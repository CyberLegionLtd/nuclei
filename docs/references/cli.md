# CLI Reference

Nuclei Help Menu contains the below specified options along with their purpose.

```
> nuclei --help

Nuclei is a fast, template based vulnerability scanner focusing
on extensive configurability, massive extensibility and ease of use.

Usage:
  nuclei [flags]

Flags:
TARGET:
   -u, -target string[]  target URLs/hosts to scan
   -l, -list string      path to file containing a list of target URLs/hosts to scan (one per line)
   -resume               Resume scan using resume.cfg (clustering will be disabled)

TEMPLATES:
   -t, -templates string[]      template or template directory paths to include in the scan
   -tu, -template-url string[]  URL containing list of templates to run
   -nt, -new-templates          run only new templates added in latest nuclei-templates release
   -w, -workflows string[]      workflow or workflow directory paths to include in the scan
   -wu, -workflow-url string[]  URL containing list of workflows to run
   -validate                    validate the passed templates to nuclei
   -tl                          list all available templates

FILTERING:
   -tags string[]                    execute a subset of templates that contain the provided tags
   -itags, -include-tags string[]    tags from the default deny list that permit executing more intrusive templates
   -etags, -exclude-tags string[]    exclude templates with the provided tags
   -it, -include-templates string[]  templates to be executed even if they are excluded either by default or configuration
   -et, -exclude-templates string[]  template or template directory paths to exclude
   -s, -severity value[]             Templates to run based on severity. Possible values: info, low, medium, high, critical
   -es, -exclude-severity value[]    Templates to exclude based on severity. Possible values: info, low, medium, high, critical
   -pt, -type value[]                protocol types to be executed. Possible values: dns, file, http, headless, network, workflow, ssl, websocket, whois
   -ept, -exclude-type value[]       protocol types to not be executed. Possible values: dns, file, http, headless, network, workflow, ssl, websocket, whois
   -a, -author string[]              execute templates that are (co-)created by the specified authors

OUTPUT:
   -o, -output string            output file to write found issues/vulnerabilities
   -silent                       display findings only
   -nc, -no-color                disable output content coloring (ANSI escape codes)
   -json                         write output in JSONL(ines) format
   -irr, -include-rr             include request/response pairs in the JSONL output (for findings only)
   -nm, -no-meta                 don't display match metadata
   -nts, -no-timestamp           don't display timestamp metadata in CLI output
   -rdb, -report-db string       local nuclei reporting database (always use this to persist report data)
   -ms, -matcher-status          show optional match failure status
   -me, -markdown-export string  directory to export results in markdown format
   -se, -sarif-export string     file to export results in SARIF format

CONFIGURATIONS:
   -config string              path to the nuclei configuration file
   -rc, -report-config string  nuclei reporting module configuration file
   -H, -header string[]        custom headers in header:value format
   -V, -var value              custom vars in var=value format
   -r, -resolvers string       file containing resolver list for nuclei
   -sr, -system-resolvers      use system DNS resolving as error fallback
   -passive                    enable passive HTTP response processing mode
   -ev, -env-vars              enable environment variables to be used in template
   -cc, -client-cert string    client certificate file (PEM-encoded) used for authenticating against scanned hosts
   -ck, -client-key string     client key file (PEM-encoded) used for authenticating against scanned hosts
   -ca, -client-ca string      client certificate authority file (PEM-encoded) used for authenticating against scanned hosts

INTERACTSH:
   -iserver, -interactsh-server string  interactsh server url for self-hosted instance (default "https://interact.sh")
   -itoken, -interactsh-token string    authentication token for self-hosted interactsh server
   -interactions-cache-size int         number of requests to keep in the interactions cache (default 5000)
   -interactions-eviction int           number of seconds to wait before evicting requests from cache (default 60)
   -interactions-poll-duration int      number of seconds to wait before each interaction poll request (default 5)
   -interactions-cooldown-period int    extra time for interaction polling before exiting (default 5)
   -ni, -no-interactsh                  disable interactsh server for OAST testing, exclude OAST based templates

RATE-LIMIT:
   -rl, -rate-limit int            maximum number of requests to send per second (default 150)
   -rlm, -rate-limit-minute int    maximum number of requests to send per minute
   -bs, -bulk-size int             maximum number of hosts to be analyzed in parallel per template (default 25)
   -c, -concurrency int            maximum number of templates to be executed in parallel (default 25)
   -hbs, -headless-bulk-size int   maximum number of headless hosts to be analyzed in parallel per template (default 10)
   -hc, -headless-concurrency int  maximum number of headless templates to be executed in parallel (default 10)

OPTIMIZATIONS:
   -timeout int               time to wait in seconds before timeout (default 5)
   -retries int               number of times to retry a failed request (default 1)
   -mhe, -max-host-error int  max errors for a host before skipping from scan (default 30)
   -project                   use a project folder to avoid sending same request multiple times
   -project-path string       set a specific project path (default "/var/folders/wd/xgh2wd216cv71wf0bqtfq9r00000gn/T/")
   -spm, -stop-at-first-path  stop processing HTTP requests after the first match (may break template/workflow logic)
   -stream                    Stream mode - start elaborating without sorting the input

HEADLESS:
   -headless            enable templates that require headless browser support (root user on linux will disable sandbox)
   -page-timeout int    seconds to wait for each page in headless mode (default 20)
   -sb, -show-browser   show the browser on the screen when running templates with headless mode
   -sc, -system-chrome  Use local installed chrome browser instead of nuclei installed

DEBUG:
   -debug                    show all requests and responses
   -debug-req                show all sent requests
   -debug-resp               show all received responses
   -p, -proxy string[]       List of HTTP(s)/SOCKS5 proxy to use (comma separated or file input)
   -tlog, -trace-log string  file to write sent requests trace log
   -elog, -error-log string  file to write sent requests error log
   -version                  show nuclei version
   -v, -verbose              show verbose output
   -vv                       display templates loaded for scan
   -tv, -templates-version   shows the version of the installed nuclei-templates

UPDATE:
   -update                        update nuclei engine to the latest released version
   -ut, -update-templates         update nuclei-templates to latest released version
   -ud, -update-directory string  overwrite the default directory to install nuclei-templates (default "/Users/ice3man/nuclei-templates")
   -duc, -disable-update-check    disable automatic nuclei/templates update check

STATISTICS:
   -stats                    display statistics about the running scan
   -sj, -stats-json          write statistics data to an output file in JSONL(ines) format
   -si, -stats-interval int  number of seconds to wait between showing a statistics update (default 5)
   -m, -metrics              expose nuclei metrics on a port
   -mp, -metrics-port int    port to expose nuclei metrics on (default 9092)
```