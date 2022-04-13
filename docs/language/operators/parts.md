# Part Specification

The following parts can be matched on for different protocols in Nuclei using matchers and extractors. A detailed list of part for each protocol is provided below.

## DNS

| Part          | Description of the part                              |
|---------------|------------------------------------------------------|
| template-id   | ID of the template executed                          |
| template-info | Info Block of the template executed                  |
| template-path | Path of the template executed                        |
| host          | Host is the input to the template                    |
| matched       | Matched is the input which was matched upon          |
| request       | Request contains the DNS request in text format      |
| type          | Type is the type of request made                     |
| rcode         | Rcode field returned for the DNS request             |
| question      | Question contains the DNS question field             |
| extra         | Extra contains the DNS response extra field          |
| answer        | Answer contains the DNS response answer field        |
| ns            | NS contains the DNS response NS field                |
| raw,body,all  | Raw contains the raw DNS response (default)          |
| trace         | Trace contains trace data for DNS request if enabled |


## File

| Part              | Description of the part                      |
|-------------------|----------------------------------------------|
| template-id       | ID of the template executed                  |
| template-info     | Info Block of the template executed          |
| template-path     | Path of the template executed                |
| matched           | Matched is the input which was matched upon  |
| path              | Path is the path of file on local filesystem |
| type              | Type is the type of request made             |
| raw,body,all,data | Raw contains the raw file contents           |