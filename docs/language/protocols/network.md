# Network

Nuclei can act as an automatable **Netcat**, allowing users to send bytes across the wire and receive them, while providing matching and extracting capabilities on the response.

Network Requests start with a `network` block which specifies the start of the requests for the template.

## host

Host to send network requests to. Dynamic variables can be placed in the path to modify its value on runtime. The following variables are supported in network host field.

| Variable | Description of variable | Example | 
|----------|-------------------------|---------|
| Host  | Host name from Input | localhost | 
| Port  | Port from Input | 443 | 
| Hostname | Host from Input | localhost:443 | 

```yaml
# Example host field
host: "{{Hostname}}"
```

Nuclei can also do TLS connection to the target server. Just add `tls://` as prefix before the Hostname and you're good to go.

```yaml
# TLS connection host field
host:
  - "tls://{{Hostname}}"
```

If a port is specified in the host field, the user supplied port is ignored and the template port takes precedence.

## attack & payloads

Payloads can be used to specify a list payloads values or files along with the attack type which can be `battering ram`, `clusterbomb` or `pitchfork`. 

```yaml
# attack & payloads for a password bruteforce clusterbomb scenario
attack: clusterbomb
payloads:
  username:
    - admin
    - root
  password:
    - password
    - toor
    - guest
inputs:
  - data: "USER {{username}}\r\nPASS {{password}}\r\n"
```

## inputs

Inputs are the data that will be sent to the server, and optionally any data to read from the server.

At it's most simple, just specify a string, and it will be sent across the network socket.

```yaml
# inputs is the list of inputs to send to the server
inputs: 
  - data: "TEST\r\n"
```

You can also send hex encoded text that will be first decoded and the raw bytes will be sent to the server.

```yaml
inputs:
  - data: "50494e47"
    type: hex
  - data: "\r\n"
```

Helper function expressions can also be defined in input and will be first evaluated and then sent to the server. The last Hex Encoded example can be sent with helper functions this way -

```yaml
inputs:
  - data: 'hex_decode("50494e47")\r\n'
```

One last thing that can be done with inputs is reading data from the socket. Specifying read-size with a non-zero value will do the trick. You can also assign the read data some name, so matching can be done on that part.

```yaml
inputs:
  - read-size: 8
```

Example with reading a number of bytes, and only matching on them.

```yaml
inputs:
  - read-size: 8
    name: prefix
...
matchers:
  - type: word
    part: prefix
    words: 
      - "CAFEBABE"
```

Multiple steps can be chained together in sequence to do network reading / writing.

## read-size

Number of bytes to read from the network socket at the end of the connection. By default, 1024 bytes are read from the socket.

```yaml
# read 8 bytes from socket
read-size: 8
```

## read-all

Read-all keeps reading data from the network socket until EOF is reached. It can be used when the size of the server response is not known beforehand.

```yaml
# keep reading from socket until we reach end
read-all: true
```

## Example

The final example template file for a hex encoded input to detect **MongoDB** running on servers with working matchers is provided below.

```yaml
id: input-expressions-mongodb-detect

info:
  name: Input Expression MongoDB Detection
  author: pd-team
  severity: info
  reference: https://github.com/orleven/Tentacle

network:
  - inputs:
      - data: "{{hex_decode('3a000000a741000000000000d40700000000000061646d696e2e24636d640000000000ffffffff130000001069736d6173746572000100000000')}}"
    host:
      - "{{Hostname}}"
    read-size: 2048
    matchers:
      - type: word
        words:
          - "logicalSessionTimeout"
          - "localTime"
```