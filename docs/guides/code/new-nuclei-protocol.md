# Adding A New Protocol To Nuclei

Protocols form the core of Nuclei Engine. All the request types like `http`, `dns`, etc. are implemented in form of protocol requests.

A protocol must implement the `Protocol` and `Request` interfaces described above in `pkg/protocols`. We'll take the example of an existing protocol implementation - websocket for this short reference around Nuclei internals.

The code for the websocket protocol is contained in `pkg/protocols/others/websocket`. 

Below a high level skeleton of the websocket implementation is provided with all the important parts present.

```go
package websocket

// Request is a request for the Websocket protocol
type Request struct {
	// Operators for the current request go here.
	operators.Operators `yaml:",inline,omitempty"`
	CompiledOperators   *operators.Operators `yaml:"-"`

	// description: |
	//   Address contains address for the request
	Address string `yaml:"address,omitempty" jsonschema:"title=address for the websocket request,description=Address contains address for the request"`

    // declarations here
}

// Compile compiles the request generators preparing any requests possible.
func (r *Request) Compile(options *protocols.ExecuterOptions) error {
	r.options = options

    // request compilation here as well as client creation
 
	if len(r.Matchers) > 0 || len(r.Extractors) > 0 {
		compiled := &r.Operators
		if err := compiled.Compile(); err != nil {
			return errors.Wrap(err, "could not compile operators")
		}
		r.CompiledOperators = compiled
	}
	return nil
}

// Requests returns the total number of requests the rule will perform
func (r *Request) Requests() int {
	if r.generator != nil {
		return r.generator.NewIterator().Total()
	}
	return 1
}

// GetID returns the ID for the request if any.
func (r *Request) GetID() string {
	return ""
}

// ExecuteWithResults executes the protocol requests and returns results instead of writing them.
func (r *Request) ExecuteWithResults(input string, dynamicValues, previous output.InternalEvent, callback protocols.OutputEventCallback) error {
    // payloads init here
	if err := r.executeRequestWithPayloads(input, hostname, value, previous, callback); err != nil {
		return err
	}
	return nil
}

// ExecuteWithResults executes the protocol requests and returns results instead of writing them.
func (r *Request) executeRequestWithPayloads(input, hostname string, dynamicValues, previous output.InternalEvent, callback protocols.OutputEventCallback) error {
	header := http.Header{}

    // make the actual request here after setting all options

	event := eventcreator.CreateEventWithAdditionalOptions(r, data, r.options.Options.Debug || r.options.Options.DebugResponse, func(internalWrappedEvent *output.InternalWrappedEvent) {
		internalWrappedEvent.OperatorsResult.PayloadValues = payloadValues
	})
	if r.options.Options.Debug || r.options.Options.DebugResponse {
		responseOutput := responseBuilder.String()
		gologger.Debug().Msgf("[%s] Dumped Websocket response for %s", r.options.TemplateID, input)
		gologger.Print().Msgf("%s", responsehighlighter.Highlight(event.OperatorsResult, responseOutput, r.options.Options.NoColor))
	}

	callback(event)
	return nil
}

func (r *Request) MakeResultEventItem(wrapped *output.InternalWrappedEvent) *output.ResultEvent {
	data := &output.ResultEvent{
		TemplateID:       types.ToString(r.options.TemplateID),
		TemplatePath:     types.ToString(r.options.TemplatePath),
		// ... setting more values for result event
	}
	return data
}

// Match performs matching operation for a matcher on model and returns:
// true and a list of matched snippets if the matcher type is supports it
// otherwise false and an empty string slice
func (r *Request) Match(data map[string]interface{}, matcher *matchers.Matcher) (bool, []string) {
	return protocols.MakeDefaultMatchFunc(data, matcher)
}

// Extract performs extracting operation for an extractor on model and returns true or false.
func (r *Request) Extract(data map[string]interface{}, matcher *extractors.Extractor) map[string]struct{} {
	return protocols.MakeDefaultExtractFunc(data, matcher)
}

// MakeResultEvent creates a result event from internal wrapped event
func (r *Request) MakeResultEvent(wrapped *output.InternalWrappedEvent) []*output.ResultEvent {
	return protocols.MakeDefaultResultEvent(r, wrapped)
}

// GetCompiledOperators returns a list of the compiled operators
func (r *Request) GetCompiledOperators() []*operators.Operators {
	return []*operators.Operators{r.CompiledOperators}
}

// Type returns the type of the protocol request
func (r *Request) Type() templateTypes.ProtocolType {
	return templateTypes.WebsocketProtocol
}
```

Almost all of these protocols have boilerplate functions for which default implementations have been provided in the `providers` package. Examples are the implementation of `Match`, `Extract`, `MakeResultEvent`, GetCompiledOperators`, etc. which are almost same throughout Nuclei protocols code. It is enough to copy-paste them unless customization is required.

`eventcreator` package offers `CreateEventWithAdditionalOptions` function which can be used to create result events after doing request execution.

Step by step description of how to add a new protocol to Nuclei - 

1. Add the protocol implementation in `pkg/protocols` directory. If it's a small protocol with fewer options, considering adding it to the `pkg/protocols/others` directory. Add the enum for the new protocol to `v2/pkg/templates/types/types.go`.

2. Add the protocol request structure to the `Template` structure fields. This is done in `pkg/templates/templates.go` with the corresponding import line.

```go

import (
	...
	"github.com/projectdiscovery/nuclei/v2/pkg/protocols/others/websocket"
)

// Template is a YAML input file which defines all the requests and
// other metadata for a template.
type Template struct {
	...
	// description: |
	//   Websocket contains the Websocket request to make in the template.
	RequestsWebsocket []*websocket.Request `yaml:"websocket,omitempty" json:"websocket,omitempty" jsonschema:"title=websocket requests to make,description=Websocket requests to make for the template"`
	...
}
```

Also add the protocol case to the `Type` function as well as the `TemplateTypes` array in the same `templates.go` file.

```go
// TemplateTypes is a list of accepted template types
var TemplateTypes = []string{
	...
	"websocket",
}

// Type returns the type of the template
func (t *Template) Type() templateTypes.ProtocolType {
	...
	case len(t.RequestsWebsocket) > 0:
		return templateTypes.WebsocketProtocol
	default:
		return ""
	}
}
```

3. Add the protocol request to the `Requests` function and `compileProtocolRequests` function in the `compile.go` file in same directory.

```go

// Requests return the total request count for the template
func (template *Template) Requests() int {
	return len(template.RequestsDNS) +
		...
		len(template.RequestsSSL) +
		len(template.RequestsWebsocket)
}


// compileProtocolRequests compiles all the protocol requests for the template
func (template *Template) compileProtocolRequests(options protocols.ExecuterOptions) error {
	...

	case len(template.RequestsWebsocket) > 0:
		requests = template.convertRequestToProtocolsRequest(template.RequestsWebsocket)
	}
	template.Executer = executer.NewExecuter(requests, &options)
	return nil
}
```

That's it, you've added a new protocol to Nuclei. The next good step would be to write integration tests which are described in `integration-tests` and `cmd/integration-tests` directories.
