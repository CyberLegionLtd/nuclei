<h1 align="center">
  <br>
  <a href="https://nuclei.projectdiscovery.io"><img src="static/nuclei-logo.png" width="200px" alt="Nuclei"></a>
</h1>

<h4 align="center">Fast and customisable vulnerability scanner based on simple YAML based DSL.</h4>


<p align="center">
<img src="https://img.shields.io/github/go-mod/go-version/projectdiscovery/nuclei?filename=v2%2Fgo.mod">
<a href="https://github.com/projectdiscovery/nuclei/releases"><img src="https://img.shields.io/github/downloads/projectdiscovery/nuclei/total">
<a href="https://github.com/projectdiscovery/nuclei/graphs/contributors"><img src="https://img.shields.io/github/contributors-anon/projectdiscovery/nuclei">
<a href="https://github.com/projectdiscovery/nuclei/releases/"><img src="https://img.shields.io/github/release/projectdiscovery/nuclei">
<a href="https://github.com/projectdiscovery/nuclei/issues"><img src="https://img.shields.io/github/issues-raw/projectdiscovery/nuclei">
<a href="https://github.com/projectdiscovery/nuclei/discussions"><img src="https://img.shields.io/github/discussions/projectdiscovery/nuclei">
<a href="https://discord.gg/projectdiscovery"><img src="https://img.shields.io/discord/695645237418131507.svg?logo=discord"></a>
<a href="https://twitter.com/pdnuclei"><img src="https://img.shields.io/twitter/follow/pdnuclei.svg?logo=twitter"></a>
</p>
      
<p align="center">
  <a href="#how-it-works">How</a> •
  <a href="#install-nuclei">Install</a> •
  <a href="#for-security-engineers">For Security Engineers</a> •
  <a href="#for-developers-and-organisations">For Developers</a> •
  <a href="https://nuclei.projectdiscovery.io/nuclei/get-started/">Documentation</a> •
  <a href="#credits">Credits</a> •
  <a href="https://nuclei.projectdiscovery.io/faq/nuclei/">FAQs</a> •
  <a href="https://discord.gg/projectdiscovery">Join Discord</a>
</p>

<p align="center">
  <a href="https://github.com/projectdiscovery/nuclei/blob/master/README.md">English</a> •
  <a href="https://github.com/projectdiscovery/nuclei/blob/master/README_CN.md">中文</a> •
  <a href="https://github.com/projectdiscovery/nuclei/blob/master/README_KR.md">Korean</a>
</p>

---

Nuclei is used to send requests across targets based on a template, leading to zero false positives and providing fast scanning on a large number of hosts. Nuclei offers scanning for a variety of protocols, including TCP, DNS, HTTP, SSL, File, Whois, Websocket, Headless etc. With powerful and flexible templating, Nuclei can be used to model all kinds of security checks.

We have a [dedicated repository](https://github.com/projectdiscovery/nuclei-templates) that houses various type of vulnerability templates contributed by **more than 300** security researchers and engineers.



## How it works


<h3 align="center">
  <img src="static/nuclei-flow.jpg" alt="nuclei-flow" width="700px"></a>
</h3>


# Install Nuclei

Nuclei requires **go1.18** to install successfully. Run the following command to install the latest version -

```sh
go install -v github.com/projectdiscovery/nuclei/v2/cmd/nuclei@latest
```

**More installation [methods can be found here](https://nuclei.projectdiscovery.io/nuclei/get-started/).**

<table>
<tr>
<td>  

### Nuclei Templates

Nuclei has built-in support for automatic template download/update as default since version [v2.5.2](https://github.com/projectdiscovery/nuclei/releases/tag/v2.5.2). [**Nuclei-Templates**](https://github.com/projectdiscovery/nuclei-templates) project provides a community-contributed list of ready-to-use templates that is constantly updated.

You may still use the `update-templates` flag to update the nuclei templates at any time; You can write your own checks for your individual workflow and needs following Nuclei's [templating guide](https://nuclei.projectdiscovery.io/templating-guide/).

The YAML DSL reference syntax is available [here](SYNTAX-REFERENCE.md).

</td>
</tr>
</table>

### Usage

```sh
nuclei -h
```

This will display help for the tool. Here are all the switches it supports.


```console
Nuclei is a fast, template based vulnerability scanner focusing
on extensive configurability, massive extensibility and ease of use.

Usage:
  nuclei [flags]

Flags:
TARGET:
   -u, -target string[]  target URLs/hosts to scan
   -l, -list string      path to file containing a list of target URLs/hosts to scan (one per line)
   -resume string        Resume scan using resume.cfg (clustering will be disabled)

TEMPLATES:
   -nt, -new-templates                    run only new templates added in latest nuclei-templates release
   -ntv, -new-templates-version string[]  run new templates added in specific version
   -as, -automatic-scan                   automatic web scan using wappalyzer technology detection to tags mapping
   -t, -templates string[]                list of template or template directory to run (comma-separated, file)
   -tu, -template-url string[]            list of template urls to run (comma-separated, file)
   -w, -workflows string[]                list of workflow or workflow directory to run (comma-separated, file)
   -wu, -workflow-url string[]            list of workflow urls to run (comma-separated, file)
   -validate                              validate the passed templates to nuclei
   -nss, -no-strict-syntax                Disable strict syntax check on templates
   -tl                                    list all available templates

FILTERING:
   -a, -author string[]               templates to run based on authors (comma-separated, file)
   -tags string[]                     templates to run based on tags (comma-separated, file)
   -etags, -exclude-tags string[]     templates to exclude based on tags (comma-separated, file)
   -itags, -include-tags string[]     tags to be executed even if they are excluded either by default or configuration
   -id, -template-id string[]         templates to run based on template ids (comma-separated, file)
   -eid, -exclude-id string[]         templates to exclude based on template ids (comma-separated, file)
   -it, -include-templates string[]   templates to be executed even if they are excluded either by default or configuration
   -et, -exclude-templates string[]   template or template directory to exclude (comma-separated, file)
   -em, -exclude-matchers string[]    template matchers to exclude in result
   -s, -severity value[]              templates to run based on severity. Possible values: info, low, medium, high, critical, unknown
   -es, -exclude-severity value[]     templates to exclude based on severity. Possible values: info, low, medium, high, critical, unknown
   -pt, -type value[]                 templates to run based on protocol type. Possible values: dns, file, http, headless, network, workflow, ssl, websocket, whois
   -ept, -exclude-type value[]        templates to exclude based on protocol type. Possible values: dns, file, http, headless, network, workflow, ssl, websocket, whois
   -tc, -template-condition string[]  templates to run based on expression condition

OUTPUT:
   -o, -output string            output file to write found issues/vulnerabilities
   -sresp, -store-resp           store all request/response passed through nuclei to output directory
   -srd, -store-resp-dir string  store all request/response passed through nuclei to custom directory (default "output")
   -silent                       display findings only
   -nc, -no-color                disable output content coloring (ANSI escape codes)
   -json                         write output in JSONL(ines) format
   -irr, -include-rr             include request/response pairs in the JSONL output (for findings only)
   -nm, -no-meta                 disable printing result metadata in cli output
   -nts, -no-timestamp           disable printing timestamp in cli output
   -rdb, -report-db string       nuclei reporting database (always use this to persist report data)
   -ms, -matcher-status          display match failure status
   -me, -markdown-export string  directory to export results in markdown format
   -se, -sarif-export string     file to export results in SARIF format

CONFIGURATIONS:
   -config string              path to the nuclei configuration file
   -fr, -follow-redirects      enable following redirects for http templates
   -mr, -max-redirects int     max number of redirects to follow for http templates (default 10)
   -dr, -disable-redirects     disable redirects for http templates
   -rc, -report-config string  nuclei reporting module configuration file
   -H, -header string[]        custom header/cookie to include in all http request in header:value format (cli, file)
   -V, -var value              custom vars in key=value format
   -r, -resolvers string       file containing resolver list for nuclei
   -sr, -system-resolvers      use system DNS resolving as error fallback
   -passive                    enable passive HTTP response processing mode
   -ev, -env-vars              enable environment variables to be used in template
   -cc, -client-cert string    client certificate file (PEM-encoded) used for authenticating against scanned hosts
   -ck, -client-key string     client key file (PEM-encoded) used for authenticating against scanned hosts
   -ca, -client-ca string      client certificate authority file (PEM-encoded) used for authenticating against scanned hosts
   -sml, -show-match-line      show match lines for file templates, works with extractors only
   -ztls                       use ztls library with autofallback to standard one for tls13
   -sni string                 tls sni hostname to use (default: input domain name)
   -i, -interface string       network interface to use for network scan
   -sip, -source-ip string     source ip address to use for network scan
   -config-directory string    Override the default config path ($home/.config)

INTERACTSH:
   -iserver, -interactsh-server string  interactsh server url for self-hosted instance (default: oast.pro,oast.live,oast.site,oast.online,oast.fun,oast.me)
   -itoken, -interactsh-token string    authentication token for self-hosted interactsh server
   -interactions-cache-size int         number of requests to keep in the interactions cache (default 5000)
   -interactions-eviction int           number of seconds to wait before evicting requests from cache (default 60)
   -interactions-poll-duration int      number of seconds to wait before each interaction poll request (default 5)
   -interactions-cooldown-period int    extra time for interaction polling before exiting (default 5)
   -ni, -no-interactsh                  disable interactsh server for OAST testing, exclude OAST based templates

RATE-LIMIT:
   -rl, -rate-limit int               maximum number of requests to send per second (default 150)
   -rlm, -rate-limit-minute int       maximum number of requests to send per minute
   -bs, -bulk-size int                maximum number of hosts to be analyzed in parallel per template (default 25)
   -c, -concurrency int               maximum number of templates to be executed in parallel (default 25)
   -hbs, -headless-bulk-size int      maximum number of headless hosts to be analyzed in parallel per template (default 10)
   -headc, -headless-concurrency int  maximum number of headless templates to be executed in parallel (default 10)

OPTIMIZATIONS:
   -timeout int                        time to wait in seconds before timeout (default 5)
   -retries int                        number of times to retry a failed request (default 1)
   -ldp, -leave-default-ports          leave default HTTP/HTTPS ports (eg. host:80,host:443
   -mhe, -max-host-error int           max errors for a host before skipping from scan (default 30)
   -project                            use a project folder to avoid sending same request multiple times
   -project-path string                set a specific project path
   -spm, -stop-at-first-path           stop processing HTTP requests after the first match (may break template/workflow logic)
   -stream                             stream mode - start elaborating without sorting the input
   -irt, -input-read-timeout duration  timeout on input read (default 3m0s)
   -no-stdin                           Disable Stdin processing

HEADLESS:
   -headless                    enable templates that require headless browser support (root user on linux will disable sandbox)
   -page-timeout int            seconds to wait for each page in headless mode (default 20)
   -sb, -show-browser           show the browser on the screen when running templates with headless mode
   -sc, -system-chrome          Use local installed chrome browser instead of nuclei installed
   -lha, -list-headless-action  list available headless actions

DEBUG:
   -debug                        show all requests and responses
   -dreq, -debug-req             show all sent requests
   -dresp, -debug-resp           show all received responses
   -p, -proxy string[]           list of http/socks5 proxy to use (comma separated or file input)
   -pi, -proxy-internal          proxy all internal requests
   -ldf, -list-dsl-function      list all supported DSL function signatures
   -tlog, -trace-log string      file to write sent requests trace log
   -elog, -error-log string      file to write sent requests error log
   -version                      show nuclei version
   -hm, -hang-monitor            enable nuclei hang monitoring
   -v, -verbose                  show verbose output
   -profile-mem string           optional nuclei memory profile dump file
   -vv                           display templates loaded for scan
   -ep, -enable-pprof            enable pprof debugging server
   -tv, -templates-version       shows the version of the installed nuclei-templates
   -hc, -health-check            run diagnostic check up

UPDATE:
   -update                        update nuclei engine to the latest released version
   -ut, -update-templates         update nuclei-templates to latest released version
   -ud, -update-directory string  overwrite the default directory to install nuclei-templates
   -duc, -disable-update-check    disable automatic nuclei/templates update check

STATISTICS:
   -stats                    display statistics about the running scan
   -sj, -stats-json          write statistics data to an output file in JSONL(ines) format
   -si, -stats-interval int  number of seconds to wait between showing a statistics update (default 5)
   -m, -metrics              expose nuclei metrics on a port
   -mp, -metrics-port int    port to expose nuclei metrics on (default 9092)
```

### Running Nuclei

Scanning target domain with [community-curated](https://github.com/projectdiscovery/nuclei-templates) nuclei templates.

```sh
nuclei -u https://example.com
```

Scanning target URLs with [community-curated](https://github.com/projectdiscovery/nuclei-templates) nuclei templates.

```sh
nuclei -list urls.txt
```

Example of `urls.txt`:

```yaml
http://example.com
http://app.example.com
http://test.example.com
http://uat.example.com
```

**More detailed examples of running nuclei can be found [here](https://nuclei.projectdiscovery.io/nuclei/get-started/#running-nuclei).**

# For Security Engineers

Nuclei offers great number of features that are helpful for security engineers to customise workflow in their organisation. With the varieties of scan capabilities (like DNS, HTTP, TCP), security engineers can easily create their suite of custom checks with Nuclei.

- Varieties of protocols supported: TCP, DNS, HTTP, File, etc
- Achieve complex vulnerability steps with workflows and [dynamic requests.](https://blog.projectdiscovery.io/nuclei-unleashed-quickly-write-complex-exploits/)
- Easy to integrate into CI/CD, designed to be easily integrated into regression cycle to actively check the fix and re-appearance of vulnerability. 

<h1 align="left">
  <a href="https://nuclei.projectdiscovery.io/nuclei/get-started/"><img src="static/learn-more-button.png" width="170px" alt="Learn More"></a>
</h1>

<table>
<tr>
<td>  

**For Bug Bounty hunters:**

Nuclei allows you to customise your testing approach with your own suite of checks and easily run across your bug bounty programs. Moreover, Nuclei can be easily integrated into any continuous scanning workflow.

- Designed to be easily integrated into other tool workflow.
- Can process thousands of hosts in few minutes.
- Easily automate your custom testing approach with our simple YAML DSL.

Please check our other open-source projects that might fit into your bug bounty workflow: [github.com/projectdiscovery](http://github.com/projectdiscovery), we also host daily [refresh of DNS data at Chaos](http://chaos.projectdiscovery.io).

</td>
</tr>
</table>

<table>
<tr>
<td>
  
**For Penetration Testers:**

Nuclei immensely improve how you approach security assessment by augmenting the manual, repetitive processes. Consultancies are already converting their manual assessment steps with Nuclei, it allows them to run set of their custom assessment approach across thousands of hosts in an automated manner. 

Pen-testers get the full power of our public templates and customization capabilities to speed up their assessment process, and specifically with the regression cycle where you can easily verify the fix.

- Easily create your compliance, standards suite (e.g. OWASP Top 10) checklist.
- With capabilities like [fuzz](https://nuclei.projectdiscovery.io/templating-guide/#advance-fuzzing) and [workflows](https://nuclei.projectdiscovery.io/templating-guide/#workflows), complex manual steps and repetitive assessment can be easily automated with Nuclei.
- Easy to re-test vulnerability-fix by just re-running the template.

</td>
</tr>
</table>


# For Developers and Organisations

Nuclei is built with simplicity in mind, with the community backed templates by hundreds of security researchers, it allows you to stay updated with the latest security threats using continuous Nuclei scanning on the hosts. It is designed to be easily integrated into regression tests cycle, to verify the fixes and eliminate vulnerabilities from occurring in the future.

- **CI/CD:** Engineers are already utilising Nuclei within their CI/CD pipeline, it allows them to constantly monitor their staging and production environments with customised templates.
- **Continuous Regression Cycle:** With Nuclei, you can create your custom template on every new identified vulnerability and put into Nuclei engine to eliminate in the continuous regression cycle.

We have [a discussion thread around this](https://github.com/projectdiscovery/nuclei-templates/discussions/693), there are already some bug bounty programs giving incentives to hackers on writing nuclei templates with every submission, that helps them to eliminate the vulnerability across all their assets, as well as to eliminate future risk in reappearing on productions. If you're interested in implementing it in your organisation, feel free to [reach out to us](mailto:contact@projectdiscovery.io). We will be more than happy to help you in the getting started process, or you can also post into the [discussion thread for any help](https://github.com/projectdiscovery/nuclei-templates/discussions/693).

<h3 align="center">
  <img src="static/regression-with-nuclei.jpg" alt="regression-cycle-with-nuclei" width="1100px"></a>
</h3>

<h1 align="left">
  <a href="https://github.com/projectdiscovery/nuclei-action"><img src="static/learn-more-button.png" width="170px" alt="Learn More"></a>
</h1>

### Resources

- [Finding bugs with Nuclei with PinkDraconian (Robbe Van Roey)](https://www.youtube.com/watch?v=ewP0xVPW-Pk) by **[@PinkDraconian](https://twitter.com/PinkDraconian)** 
- [Nuclei: Packing a Punch with Vulnerability Scanning](https://bishopfox.com/blog/nuclei-vulnerability-scan) by **Bishopfox**
- [The WAF efficacy framework](https://www.fastly.com/blog/the-waf-efficacy-framework-measuring-the-effectiveness-of-your-waf) by **Fastly**
- [Scanning Live Web Applications with Nuclei in CI/CD Pipeline](https://blog.escape.tech/devsecops-part-iii-scanning-live-web-applications/) by **[@TristanKalos](https://twitter.com/TristanKalos)**
- [Community Powered Scanning with Nuclei](https://blog.projectdiscovery.io/community-powered-scanning-with-nuclei/)
- [Nuclei Unleashed - Quickly write complex exploits](https://blog.projectdiscovery.io/nuclei-unleashed-quickly-write-complex-exploits/)
- [Nuclei - Fuzz all the things](https://blog.projectdiscovery.io/nuclei-fuzz-all-the-things/)
- [Nuclei + Interactsh Integration for Automating OOB Testing](https://blog.projectdiscovery.io/nuclei-interactsh-integration/)
- [Weaponizes nuclei Workflows to Pwn All the Things](https://medium.com/@dwisiswant0/weaponizes-nuclei-workflows-to-pwn-all-the-things-cd01223feb77) by **[@dwisiswant0](https://github.com/dwisiswant0)**
- [How to Scan Continuously with Nuclei?](https://medium.com/@dwisiswant0/how-to-scan-continuously-with-nuclei-fcb7e9d8b8b9) by **[@dwisiswant0](https://github.com/dwisiswant0)**
- [Hack with Automation !!!](https://dhiyaneshgeek.github.io/web/security/2021/07/19/hack-with-automation/) by **[@DhiyaneshGeek](https://github.com/DhiyaneshGeek)**

### Credits

Thanks to all the amazing community [contributors for sending PRs](https://github.com/projectdiscovery/nuclei/graphs/contributors). Do also check out the below similar open-source projects that may fit in your workflow:

[FFuF](https://github.com/ffuf/ffuf), [Qsfuzz](https://github.com/ameenmaali/qsfuzz), [Inception](https://github.com/proabiral/inception), [Snallygaster](https://github.com/hannob/snallygaster), [Gofingerprint](https://github.com/Static-Flow/gofingerprint), [Sn1per](https://github.com/1N3/Sn1per/tree/master/templates), [Google tsunami](https://github.com/google/tsunami-security-scanner), [Jaeles](https://github.com/jaeles-project/jaeles), [ChopChop](https://github.com/michelin/ChopChop)

### License

Nuclei is distributed under [MIT License](https://github.com/projectdiscovery/nuclei/blob/master/LICENSE.md)

<h1 align="left">
  <a href="https://discord.gg/projectdiscovery"><img src="static/Join-Discord.png" width="380" alt="Join Discord"></a> <a href="https://nuclei.projectdiscovery.io"><img src="static/check-nuclei-documentation.png" width="380" alt="Check Nuclei Documentation"></a>
</h1>
  <p align="center">
  <a href="#https://cyberlegion.io/bug-bounty-continue-penetration-testing/">Private Bug Bounty</a> •
  <a href="#for-security-engineers">For Security Engineers</a> •
  <a href="#for-developers-and-organisations">For Developers</a> •
  <a href="https://nuclei.projectdiscovery.io/nuclei/get-started/">Documentation</a> •
  <a href="#credits">Credits</a> •
  <a href="https://nuclei.projectdiscovery.io/faq/nuclei/">FAQs</a> •
  <a href="https://discord.gg/projectdiscovery">Join Discord</a>
</p>
  We offer security testing services such as [API Pen Testing](https://cyberlegion.io/api-security-testing);
[Private Bug Bounty](https://cyberlegion.io/bug-bounty-continue-penetration-testing)
