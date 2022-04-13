# DNS

DNS Requests start with a `dns` block which specifies the start of the requests for the DNS template.

The fields supported by DNS requests are specified with details below.

## name

Name is the name of the domain to make DNS requests to. It supports templating by using `{{var}}` syntax. The following variables are available generated from the user specified input.

| Variable | Description of the variable   | Example                 |
|----------|-------------------------------|-------------------------|
| FQDN   | Input DNS Name                | www.projectdiscovery.io |
| RDN    | Domain name of Input with TLD | projectdiscovery.io     |
| DN     | Domain Name of Input          | projectdiscovery        |
| TLD    | TLD of Input                  | io                      |
| SD     | Subdomain of Input            | www                     |

```yaml
# This value will be replaced on execution with the FQDN. 
name: {{FQDN}}
```

## type

Type of the DNS request to make. The type can be any of the following - 

- A (default)
- NS
- DS
- CNAME
- SOA
- PTR
- MX
- TXT
- AAAA
- CAA

```yaml
# Example CNAME request type
type: CNAME
```

## class

Class is the class of the DNS request to make. Usually it's enough to leave it as INET. The following class values are supported - 

- inet (default)
- csnet
- chaos
- hesiod
- none
- any

```yaml
# DNS request with class INET
class: INET
```

## retries

Retries is the maximum number of retries for the DNS request. It is recommended to use a good median value from 3 to 5 for most cases. If a request exceeds number of retries, it is considered to be failed.

```yaml
# Use a retry of 3 to 5 generally
retries: 5
```

## trace

Trace enables trace mode for DNS request. This simulates `dig +trace` command. `trace-max-recursion` field is also required while using `trace` feature.

```yaml
# An example request with DNS trace enabled
trace: true
trace-max-recursion: 10
```

The data from the trace request is available in `trace` part which can then be used by matchers and extractors.

## trace-max-recursion

This field is required along with `trace` field. It specifies the maximum number of times to recurse during DNS trace request. 

```yaml
# A trace-max-recursion field with value as 10
trace-max-recursion: 10
```

## resolvers

List of custom resolvers to use for DNS request. By default, nuclei uses a few default resolvers. However, to override and specify custom resolvers for the request, the `resolvers` field can be used.

```yaml
# Custom resolvers
resolvers:
  - "9.8.8.9:53"
  - "10.0.0.1"
```

An example DNS template that detects AWS EC2 CNAME is provided below to illustrate how a DNS template looks like.

```yaml
id: ec2-detection

info:
  name: AWS EC2 Detection
  author: melbadry9
  severity: info
  description: Amazon Elastic Compute Cloud (EC2) detected.
  tags: dns,ec2,aws
  reference:
    - https://blog.melbadry9.xyz/dangling-dns/aws/ddns-ec2-current-state
  classification:
    cvss-metrics: CVSS:3.0/AV:N/AC:L/PR:N/UI:N/S:U/C:N/I:N/A:N
    cvss-score: 0.0
    cwe-id: CWE-200

dns:
  - name: "{{FQDN}}"
    type: CNAME

    extractors:
      - type: regex
        regex:
          - "ec2-[-\\d]+\\.compute[-\\d]*\\.amazonaws\\.com"
          - "ec2-[-\\d]+\\.[\\w\\d\\-]+\\.compute[-\\d]*\\.amazonaws\\.com"
```