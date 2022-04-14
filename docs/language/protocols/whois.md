# Whois

Whois requests start with `whois` block. The fields for Whois requests are specified below.

## query

Query contains the query to make to Whois server. The following variables are supported in query field.

| Variable | Description of variable | Example | 
|----------|-------------------------|---------|
| Input  | User supplied Input | google.com:443 | 
| Host  | Host for Input | google.com | 
| Hostname | Hostname from Input | localhost:443 | 

```yaml
# Example query field
query: "{{Host}}"
```

## server

Server specifies the server to make Whois (RDAP) request to.

```yaml
# Custom Whois server
server: https://rdap.namecheap.com
```

## Example

An example template for Whois protocol.

```yaml
id: basic-whois-example

info:
  name: test template for WHOIS
  author: pdteam
  severity: info

whois:
  - query: "{{Host}}"
    server: https://rdap.namecheap.com
    extractors:
      - type: kval
        kval:
          - "expiration date"
          - "registrar"
```