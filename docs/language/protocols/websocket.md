# Websocket

Websocket requests start with `websocket` block. The fields for Websocket requests are specified below.

## address

Address contains the address to make the request to. The following variables are supported in address field.

| Variable | Description of variable | Example | 
|----------|-------------------------|---------|
| Host  | Host name from Input | localhost | 
| Scheme  | Scheme for Input | https | 
| Hostname | Host from Input | localhost:443 | 
| Path | Path from Input | /test | 

```yaml
# Example address field
address: "{{Scheme}}://{{Hostname}}{{Path}}"
```

## inputs

Inputs contains the inputs for websocket protocol.

At it's most simple, just specify a string, and it will be sent across the websocket connection.

```yaml
# inputs is the list of inputs to send to the server
inputs: 
  - data: "{"secret":"{{value}}"}"
```


Helper function expressions can also be defined in input and will be first evaluated and then sent to the server.

```yaml
inputs:
  - data: 'hex_decode("50494e47")\r\n'
```

One last thing that can be done with inputs is reading data from the websocket. Specifying name with string value will do the trick. Matching can be done on the assigned name for data..

```yaml
inputs:
  - name: response
...
matchers:
  - type: word
    words:
      - "data"
    part: response
```

## headers

Headers contains any extra HTTP headers for the websocket request. 

```yaml
# Websocket headers example
headers:
  Origin: https://some-random-origin.com
```

## attack & payload

Payloads can be used to specify a list payloads values or files along with the attack type which can be `battering ram`, `clusterbomb` or `pitchfork`. 

```yaml
# attack & payloads for a clusterbomb scenario
attack: clusterbomb
payloads:
  username:
    - admin
  password:
    - password
    - guest
```

## Example

An example of template that checks for Cross-Site-Websocket-Hijacking (CSWH) issues is provided below.

```yaml
id: cswsh

info:
  name: Cross Site Websocket Hijacking Test
  author: pdteam
  severity: low

websocket:
  - address: '{{Scheme}}://{{Hostname}}{{Path}}'
    headers: 
      Origin: 'http://evil.com'
    matchers:
      - type: word
        words:
          - true
        part: success
```