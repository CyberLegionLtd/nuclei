# Part Specification

The following parts can be matched on for different protocols in Nuclei using matchers and extractors. A detailed list of part for each protocol is provided below.

## Common

| Part          | Description of the part             |
|---------------|-------------------------------------|
| template-id   | ID of the template executed         |
| template-info | Info Block of the template executed |
| template-path | Path of the template executed       |

## DNS

| Part         | Description of the part                              |
|--------------|------------------------------------------------------|
| host         | Host is the input to the template                    |
| matched      | Matched is the input which was matched upon          |
| request      | Request contains the DNS request in text format      |
| type         | Type is the type of request made                     |
| rcode        | Rcode field returned for the DNS request             |
| question     | Question contains the DNS question field             |
| extra        | Extra contains the DNS response extra field          |
| answer       | Answer contains the DNS response answer field        |
| ns           | NS contains the DNS response NS field                |
| raw,body,all | Raw contains the raw DNS response (default)          |
| trace        | Trace contains trace data for DNS request if enabled |


## File

| Part              | Description of the part                      |
|-------------------|----------------------------------------------|
| matched           | Matched is the input which was matched upon  |
| path              | Path is the path of file on local filesystem |
| type              | Type is the type of request made             |
| raw,body,all,data | Raw contains the raw file contents           |

## Headless

| Part           | Description of the part                          |
|----------------|--------------------------------------------------|
| host           | Host is the input to the template                |
| matched        | Matched is the input which was matched upon      |
| type           | Type is the type of request made                 |
| req            | Headless request made from the client            |
| resp,body,data | Headless response received from client (default) |

## Network

| Part          | Description of the part                         |
|---------------|-------------------------------------------------|
| host          | Host is the input to the template               |
| matched       | Matched is the input which was matched upon     |
| type          | Type is the type of request made                |
| request       | Network request made from the client            |
| body,all,data | Network response received from server (default) |
| raw           | Full Network protocol data                      |

## SSL

| Part      | Description of the part                       |
|-----------|-----------------------------------------------|
| type      | Type is the type of request made              |
| response  | JSON SSL protocol handshake details           |
| not_after | Timestamp after which the remote cert expires |
| host      | Host is the input to the template             |
| matched   | Matched is the input which was matched upon   |

## Websocket

| Part     | Description of the part                                       |
|----------|---------------------------------------------------------------|
| type     | Type is the type of request made                              |
| success  | Success specifies whether websocket connection was successful |
| request  | Websocket request made to the server                          |
| response | Websocket response received from the server                   |
| host     | Host is the input to the template                             |
| matched  | Matched is the input which was matched upon                   |

## Whois

| Part     | Description of the part             |
|----------|-------------------------------------|
| type     | Type is the type of Whois request   |
| host     | Host the Whois request was made for |
| response | Whois data in JSON response format  |