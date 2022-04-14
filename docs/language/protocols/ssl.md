# SSL

SSL requests start with `ssl` block. The fields for SSL requests are specified below.

## address

Address contains the address to make SSL requests to. The following variables are supported in address field.

| Variable | Description of variable | Example | 
|----------|-------------------------|---------|
| Host  | Host name from Input | localhost | 
| Port  | Port from Input | 443 | 
| Hostname | Host from Input | localhost:443 | 


```yaml
# Example address field
address: "{{Hostname}}"
```

## min_version & max_version

min_version is the minimum TLS version required by the client. max_version is the maximum TLS version required by the client. The following values are supported for both fields.

- sslv3
- tls10
- tls11
- tls12
- tls13

```yaml
# Minimum and maximum TLS version example
min_version: sslv3
max_version: tls11
```

## cipher_suites

A list of cipher suites to request from the server. The list of cipher suites is maintained [here](https://github.com/projectdiscovery/nuclei/blob/master/v2/pkg/protocols/ssl/ciphers.go).

```yaml
# Example cipher_suite custom field
cipher_suites: 
  - TLS_AES_128_GCM_SHA256
```

## Example

An example template for detecting expired TLS version is provided below.

```yaml
id: deprecated-tls

info:
  name: Deprecated TLS Detection (inferior to TLS 1.2)
  author: righettod
  severity: info
  reference: https://ssl-config.mozilla.org/#config=intermediate
  metadata:
    shodan-query: ssl.version:sslv2 ssl.version:sslv3 ssl.version:tlsv1 ssl.version:tlsv1.1
  tags: ssl

ssl:
  - address: "{{Host}}:{{Port}}"
    min_version: sslv3
    max_version: tls11

    extractors:
      - type: json
        json:
          - " .tls_version"
```