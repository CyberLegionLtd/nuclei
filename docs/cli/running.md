# Running Nuclei

A brief guide of all the command groups nuclei supports and what they do. This will cover everything needed to fully utilise Nuclei CLI.

## Targets

Nuclei supports two ways to pass targets to perform scan on.

1. `-u` or `-target` - Accepts a single target. Can be supplied multiple times to pass multiple targets.
2. `-l` or `-list` - Accepts a file of targets as input. Can contain newline-separated target list.

Examples - 

```
nuclei -u https://example.com -u test.example.com
nuclei -l http_list.txt
```

## Templates

By default, if no templates or workflows are provided, all the templates in public templates directory are executed (except the ones in Nuclei Ignore file, more on it later).

Example - 

```bash
# This is effectively doing -> nuclei -u https://example.com -t nuclei-templates/
nuclei -u https://example.com
```

The `-t / -templates` or `-w / -workflows` flags can be used for running a list of templates or workflows. These flags support files, directories as well as Glob expressions for matching a list of files.

```bash
# Single file
nuclei -t CVE-2020-15462.yaml

# Multiple directories
nuclei -t cves/ -t exposed-panels/

# Glob pattern
nuclei -t CVE-2020-*.yaml

# Similarly for workflows
nuclei -w netsweeper-workflow.yaml -w workflows/
```

The `-nt / -new-templates` flag can be used to run all the newly added templates in the last update of the nuclei-templates repository.

## Filtering

Nuclei provides a number of filtering option to fine tune the templates being executed as per your needs. By default the filtering is applied on all public templates if no specific templates/workflows are provided with `-t/-w` flags. If provided, the user-specified path is used instead for filtering.

Nuclei contains an [Ignore File](https://github.com/projectdiscovery/nuclei-templates/blob/master/.nuclei-ignore) in the Nuclei-Templates repository. It contains some tags like `dos`, `fuzz` etc which are not executed by default. In case such templates need to be executed, `-itags / -include-tags` flag can be used.

Running templates can be customized on following metadata - 

1. Tags - Tags found in the template `info` section

```bash
# Run all templates having tag "cve"
nuclei -tags cve
# Run all templates from exposed-panels not having tag "grafana" and "kubernetes"
nuclei -t exposed-panels/ -etags "grafana,kubernetes"
# Include "dos" templates (not allowed by default) from cves directory
nuclei -t cves/ -itags "dos"
```

2. Path - Path of the template. (Relative to templates directory)

```bash
# Exclude all templates from exposed-panels/
nuclei -exclude-templates exposed-panels/
# Include an excluded directory (ex. "fuzz/")
nuclei -include-templates fuzz/
```

3. Severity - Severity of the template. low, medium high, etc

```bash
# Exclude all templates with info,low severity
nuclei -exclude-severity info,low
# Run high,critical severity vulnerabilities only
nuclei -severity high,critical
```

4. Author - Author(s) of the template

```bash
# Run templates by a specific author
nuclei -author pdteam
```

5. Type - Type of the template. dns, http, etc

```bash
# Exclude all templates of type "headless","ssl"
nuclei -exclude-type headless,ssl
# Run all templates of type "dns" and "http"
nuclei -type dns,http
```

## Output

Nuclei writes found results to the Standard Output (Stdout) and errors to the Standard Error (Stderr). `-o`/`-output` flag can be used to write the found results to a file, delimited by newlines.

`-json` flag can be used to change the mode of output to JSON. If `-irr`/`-include-rr` flag is passed, the Request / Response pair is also included in the output JSON (omitted by default to reduce size).

```bash
# Write output to a file named results.txt
nuclei -l list.txt -o results.txt

# Write JSON output to file with request response data
nuclei -l list.txt -json -irr -o results.json
```

A number of additional configuration options are provided to fine tune the Nuclei output as desired. These are specified as follows - 

1. `-silent` - Supress additional output and only display valid results
2. `-nc`,`-no-color` - Do not display colored output
3. `-nm`,`-no-meta` - Do not display match metadata
4. `-nts`,`-no-timestamp` - Do not display timestamp
5. `-ms`,`-matcher-status` - Show optional failed matchers as well in output

```bash
# Display only valid results
nuclei -u https://example.com -silent -nc -nm -nts
```

### Markdown / Sarif Export

Markdown Export support allows writing of found results in Markdown format to a folder on disk. The template used for formatting is same as one of Reporting modules.

Sarif Export allows writing of found results in Sarif file format, which is a commonly used format for Static Analysis tools.

Example - 

```bash
# Run nuclei and write markdown files to "reports" directory
nuclei -t cves/ -markdown-export "reports"

# Run nuclei and write sarif output in sarif format to "results.sarif"
nuclei -t cves/ -sarif-export "results.sarif"
```

## Configuration

### Config File

Since release of [v.2.3.2](https://blog.projectdiscovery.io/nuclei-v2-3-0-release/) nuclei uses [goflags](https://github.com/projectdiscovery/goflags) for clean CLI experience and both long/short formatted flags.

goflags comes with auto-generated config file support that coverts all available CLI flags into config file, basically you can define all CLI flags into config file to avoid repetitive CLI flags that are loaded as default for every scan of nuclei.

Default path of nuclei config file is $HOME/.config/nuclei/config.yaml, uncomment and configure the flags you wish to run by default.

Example Config - 

```yaml
# Headers to include with all HTTP request
header:
  - 'X-BugBounty-Hacker: h1/geekboy'

# Directory based template execution
templates:
  - cves/
  - vulnerabilities/
  - misconfiguration/

# Tags based template execution
tags: exposures,cve

# Template Filters
tags: exposures,cve
author: geeknik,pikpikcu,dhiyaneshdk
severity: critical,high,medium

# Template Allowlist
include-tags: dos,fuzz # Tag based inclusion (allows overwriting nuclei-ignore list)
include-templates: # Template based inclusion (allows overwriting nuclei-ignore list)
  - vulnerabilities/xxx
  - misconfiguration/xxxx

# Template Denylist
exclude-tags: info # Tag based exclusion
exclude-templates: # Template based exclusion
  - vulnerabilities/xxx
  - misconfiguration/xxxx

# Rate Limit configuration
rate-limit: 500
bulk-size: 50
concurrency: 50
```

Once configured, **config file will be used as default**, additionally custom config file can be also provided using -config flag.

```bash
# Running nuclei with custom config file
nuclei -config project.yaml -list urls.txt
```

### Reporting Configuration

Nuclei comes with reporting module supporting GitHub, GitLab, and Jira integration, this allows nuclei engine to create automatic tickets on the supported platform based on found results.

`-rc`, `-report-config` flag can be used to provide a config file to read configuration details of the platform to integrate. [Here is an example config file](../references/configuration-examples.md) for all supported platforms.

For example, to create tickets on GitHub, create a config file with the following content and replace the appropriate values:-

```yaml
# GitHub contains configuration options for GitHub issue tracker
github: 
  username: "$user"
  owner: "$user"
  token: "$token"
  project-name: "testing-project"
  issue-label: "Nuclei"
```

Running nuclei with reporting module -

```bash
# Run nuclei and use issue-tracker.yaml as reporting config
nuclei -l urls.txt -t cves/ -rc issue-tracker.yaml
```

Similarly, other platforms can be configured. Reporting module also supports basic filtering and duplicate checks to avoid duplicate ticket creation.

```yaml
# Custom allowlist allowing only high and critical
allow-list:
  severity: high,critical
```

This will ensure to only creating tickets for issues identified with high and critical severity; similarly, `deny-list` can be used to exclude issues with a specific severity.

If you are running periodic scans on the same assets, you might want to consider `-rdb`, `-report-db` flag that creates a local copy of the valid findings in the given directory utilized by reporting module to compare and create tickets for unique issues only.

```bash
# Run nuclei and use issue-tracker.yaml with local database prod for comparing results.
nuclei -l urls.txt -t cves/ -rc issue-tracker.yaml -rdb prod
```

### Custom / Environment Variables

Nuclei engine supports passing custom variables via CLI using `-V`, `-var` flag and Environment Variables by using `-ev`, `-env-vars` which exposes environment variables to the template.

These values can then be used in the template by defining them in value `{{<value>}}` placeholders.

```yaml
requests:
  - path: "{{BaseURL}}/?auth={{token}}" # value from custom variable
    headers:
      another: "{{value}}" # value from environment variable
```

Example command to expose the values - 

```bash
# Run nuclei with a custom token and environment variables
nuclei -t template.yaml -var token=value -env-vars
```

### Other Configurations

**Custom Headers**

Custom Headers can be passed by using `-H` / `-header` flag. These headers will be added to all the requests made by nuclei. 

Many BugBounty platform/programs requires you to identify the HTTP traffic you make, this can be achieved by setting custom header using config file at $HOME/.config/nuclei/config.yaml or CLI flag.

Setting custom header using config file config.yaml

```yaml
# Headers to include with each request.
header:
  - 'X-BugBounty-Hacker: h1/geekboy'
  - 'User-Agent: Mozilla/5.0 (Windows NT 10.0; WOW64) / nuclei'
```

Setting custom header using CLI flag

```bash
# Running nuclei with a custom header
nuclei -header 'User-Agent: Mozilla/5.0 (Windows NT 10.0; WOW64) / nuclei' -list urls.txt -tags cves
```

**Custom Resolvers**

By Default, nuclei uses a custom DNS resolver with **Google** and **Cloudflare** DNS servers to do the DNS resolutions. A list of custom resolvers can be passed to nuclei by using `-r`, `-resolvers` flag.

Example of a resolver list - 

```bash
9.9.9.9
9.0.0.9:54
```

```bash
# Run nuclei with a custom resolver file "resolvers.txt"
nuclei -t cves/ resolvers.txt
```

In case nuclei's DNS resolving fails, a fallback option can also be provided optionally which will use System DNS stack for doing resolutions. This behaviour can be enabled by using `-sr`, `-system-resolvers` flag.

```bash
# Run nuclei with system resolving fallback
nuclei -system-resolvers -t cves/
```

**Passive Mode**

Nuclei engine supports passive mode scanning for HTTP based template utilizing file support, with this support we can run HTTP based templates against locally stored HTTP response data collected from any other tool.

```bash
# Running nuclei passive mode against http_data folder
nuclei -passive -target http_data
```

Passive mode support is limited for templates having `{{BasedURL}}` or `{{BasedURL/}}` as base path.

**SSL Client Based Authentication**

Custom SSL Certificates, Keys and Certificate Authorities can be passed to Nuclei for authentication against hosts which require custom SSL Configuration.

`-cc`/`-client-cert`, `-ck`/`-client-key` and `-ca`/`-client-ca` can be used for passing these configurations.


## Interactsh

Nuclei comes with Interactsh integration for Out of Band (OOB) testing support. The interactsh integration can be configured by using the below provided flags.

- `-iserver`, `-interactsh-server` - Custom interactsh server
- `-itoken`, -interactsh-token` - Interactsh Token for self hosted server
- `-ni`, `-no-interactsh` - Disable interactsh server

```bash
# Run nuclei with custom interactsh server and token
nuclei -t cves/ -interactsh-server https://exfiltest.com -interactsh-token testingtokeninteractsh
```


## Rate-Limits

Nuclei comes with a number of options to control the rate-limit. This includes - the number of templates to execute concurrently, the number of hosts to process for each template concurrently and the global maxomim number of requests per minute/second that you want to send.

**Rate Limit Per Minute Or Second** - `-rl`/`-rate-limit` is the maximum number of requests to send per second. `-rlm`/`-rate-limit-minute` is the maximum number of requests to send per minute.

**Concurrency** - `-c`/`-concurrency` is the maximum number of template to process concurrently. Headless templates concurrency can be configured by using `-hc`/`-headless-concurrency` flag similarly.

**Bulk Size** - `-bs`/`-bulk-size` is the maximum number of hosts to process per template concurrently. Headless templates bulk-size can be configured by using `-hbs`/`-headless-bulk-size` flag similarly.

Feel free to play with these flags to tune your nuclei scan speed and accuracy.

```bash
# Example of nuclei rate limit configurations
nuclei -t cves/ -rate-limit 900 -c 30 -bs 30 -l list.txt
```

**Tip**

rate-limit flag takes precedence over the other two flags, the number of requests/seconds can't go beyond the value defined for rate-limit flag regardless the value of c and bulk-size flag.


## Optimizations

**Timeout and Retries**

Number of seconds to wait for a response can be configured by using `-timeout` flag. By default, a timeout of 5 seconds is used per request. The number of times a failed request is retried can be configured by using `-retries` flag. Default number of retries is 1.

```bash
# Run nuclei with timeout of 10 seconds and 3 retries
nuclei -timeout 10 -retries 3
```

**Errors Per Host**

Errors occuring per-host are kept track of across protocols. If a host crosses the maximum errors per host threshold, which is defined by `-mhe`/`-max-host-error` flag, the host is marked as not-working and discarded from the rest of the scan. By default, 30 errors per host is considered to be the threshold.

```bash
# Run nuclei with max host error threshold of 10.
nuclei -max-host-error 10
```

**Stop At First Match**

Templates include logic that decide whether on finding a match, remaining requests should be dropped or performed. This is performed by using `stop-at-first-match: true` template attribute. However, a CLI flag is provided called `-spm`/`-stop-at-first-path` which forces the remaining requests to be skipped after matches. 

```bash
# Run nuclei with stop-at-first-match globally enabled
nuclei -stop-at-first-path
```

