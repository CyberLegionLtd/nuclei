# Using Nuclei From Go Code

An example of using Nuclei From Go Code to run templates on targets is provided below.

```go
package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/logrusorgru/aurora"
	"go.uber.org/ratelimit"

	"github.com/projectdiscovery/goflags"
	"github.com/projectdiscovery/nuclei/v2/pkg/catalog"
	"github.com/projectdiscovery/nuclei/v2/pkg/catalog/config"
	"github.com/projectdiscovery/nuclei/v2/pkg/catalog/loader"
	"github.com/projectdiscovery/nuclei/v2/pkg/core"
	"github.com/projectdiscovery/nuclei/v2/pkg/core/inputs"
	"github.com/projectdiscovery/nuclei/v2/pkg/output"
	"github.com/projectdiscovery/nuclei/v2/pkg/parsers"
	"github.com/projectdiscovery/nuclei/v2/pkg/protocols"
	"github.com/projectdiscovery/nuclei/v2/pkg/protocols/common/hosterrorscache"
	"github.com/projectdiscovery/nuclei/v2/pkg/protocols/common/interactsh"
	"github.com/projectdiscovery/nuclei/v2/pkg/protocols/common/protocolinit"
	"github.com/projectdiscovery/nuclei/v2/pkg/protocols/common/protocolstate"
	"github.com/projectdiscovery/nuclei/v2/pkg/reporting"
	"github.com/projectdiscovery/nuclei/v2/pkg/testutils"
	"github.com/projectdiscovery/nuclei/v2/pkg/types"
)

func main() {
	cache := hosterrorscache.New(30, hosterrorscache.DefaultMaxHostsCount)
	defer cache.Close()

	mockProgress := &testutils.MockProgressClient{}
	reportingClient, _ := reporting.New(&reporting.Options{}, "")
	defer reportingClient.Close()

	outputWriter := testutils.NewMockOutputWriter()
	outputWriter.WriteCallback = func(event *output.ResultEvent) {
		fmt.Printf("Got Result: %v\n", event)
	}

	defaultOpts := types.DefaultOptions()
	protocolstate.Init(defaultOpts)
	protocolinit.Init(defaultOpts)

	defaultOpts.Templates = goflags.StringSlice{"dns/cname-service-detection.yaml"}
	defaultOpts.ExcludeTags = config.ReadIgnoreFile().Tags

	interactOpts := interactsh.NewDefaultOptions(outputWriter, reportingClient, mockProgress)
	interactClient, err := interactsh.New(interactOpts)
	if err != nil {
		log.Fatalf("Could not create interact client: %s\n", err)
	}
	defer interactClient.Close()

	home, _ := os.UserHomeDir()
	catalog := catalog.New(path.Join(home, "nuclei-templates"))
	executerOpts := protocols.ExecuterOptions{
		Output:          outputWriter,
		Options:         defaultOpts,
		Progress:        mockProgress,
		Catalog:         catalog,
		IssuesClient:    reportingClient,
		RateLimiter:     ratelimit.New(150),
		Interactsh:      interactClient,
		HostErrorsCache: cache,
		Colorizer:       aurora.NewAurora(true),
	}
	engine := core.New(defaultOpts)
	engine.SetExecuterOptions(executerOpts)

	workflowLoader, err := parsers.NewLoader(&executerOpts)
	if err != nil {
		log.Fatalf("Could not create workflow loader: %s\n", err)
	}
	executerOpts.WorkflowLoader = workflowLoader

	store, err := loader.New(loader.NewConfig(defaultOpts, catalog, executerOpts))
	if err != nil {
		log.Fatalf("Could not create loader client: %s\n", err)
	}
	store.Load()

	input := &inputs.SimpleInputProvider{Inputs: []string{"docs.hackerone.com"}}
	_ = engine.Execute(store.Templates(), input)
}
```