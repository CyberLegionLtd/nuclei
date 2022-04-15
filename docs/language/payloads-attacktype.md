# Payloads & Attack-type

## Payloads

Payloads contains a list of payloads for the HTTP request. The payloads can be either key-values or path to a file containing a list of values.

An example of the using payloads with local wordlist:

```yaml
# Payload fuzzing using local wordlist.
payloads:
  paths: params.txt
  header: local.txt
```

An example of the using payloads with in template wordlist support:

```yaml
# Payload fuzzing using in template wordlist.
payloads:
  password:
    - admin
    - guest
    - password
```

> Note: be careful while selecting attack type, as unexpected input will break the template.

For example, if you used clusterbomb or pitchfork as attack type and defined only one variable in the payload section, template will fail to compile, as clusterbomb or pitchfork expect more than one variable to use in the template.

## Attack-Type

Attack type is the type of payload combination to perform. The different payload types are - `batteringram`, `pitchfork` and `clusterbomb`.

Batteringram is the default type which is generally used to fuzz single parameter, clusterbomb and pitchfork for fuzzing multiple parameters which work same as classical burp intruder.

### batteringram

The battering ram attack type places the same payload value in all positions. It uses only one payload set. It loops through the payload set and replaces all positions with the payload value.

### pitchfork

The pitchfork attack type uses one payload set for each position. It places the first payload in the first position, the second payload in the second position, and so on.

It then loops through all payload sets at the same time. The first request uses the first payload from each payload set, the second request uses the second payload from each payload set, and so on.

### clusterbomb

The cluster bomb attack tries all different combinations of payloads. It still puts the first payload in the first position, and the second payload in the second position. But when it loops through the payload sets, it tries all combinations.

It then loops through all payload sets at the same time. The first request uses the first payload from each payload set, the second request uses the second payload from each payload set, and so on.

This attack type is useful for a brute-force attack. Load a list of commonly used usernames in the first payload set, and a list of commonly used passwords in the second payload set. The cluster bomb attack will then try all combinations.

More details [here](https://www.sjoerdlangkemper.nl/2017/08/02/burp-intruder-attack-types/).

An example of the using clusterbomb attack to fuzz.

```yaml
requests:
  - raw:
      - |
        POST /?file={{path}} HTTP/1.1
        User-Agent: {{header}}
        Host: {{Hostname}}

    payloads:
      path: helpers/wordlists/prams.txt
      header: helpers/wordlists/header.txt
    attack: pitchfork # Defining HTTP fuzz attack typ
```