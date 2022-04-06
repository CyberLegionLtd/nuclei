# Unique Template Matchers

The matchers block of the template is the most significant component of the template since nuclei prints results based on what we define in the matchers block. Weak matchers, in general, result in templates that produce false-positive and, in some situations, false-negative outcomes.

To write a nuclei template with unique matchers, we must keep in mind that the matcher must not only detect the specific susceptible response but also discard any random web server sending a similar response. To do so, we must consider the following requirements.

1) **Number** of matchers to use
2) **Type** of matchers to use

To write good nuclei templates, a minimum of two matchers are required. Using different types of matchers, such as HTTP **status code**, **content-type**, and **unique string** always aids in the creation of unique matchers.

| Matchers DO's              | Matchers DOn'ts                  |
| -------------------------- | ---------------------------------|
| ✅ Using matchers condition | ❌ Using single matcher          |
| ✅ Using multiple matchers  | ❌ Using only status matcher     |
| ✅ Using request condition  | ❌ Using input data as a matcher |

**Nuclei outcomes are only as excellent as their matchers**, so here's an example of an ideal matcher block: -

```yaml
   # Example matcher block
    matchers-condition: and
    matchers:

        # Status Code
      - type: status
        status:
          - 200

        # Content Type
      - type: word
        words:
          - "application/json"
        part: header

        # Response String
      - type: word
        words:
          - "Unique string from response body"
        part: body
```


## Tip from us😇

- **Nuclei** outcomes are only as excellent as **template matchers💡**
- Declare at least two matchers to reduce false positive
- Avoid matching words reflected in the URL to reduce false positive
- Avoid short word that could be encountered anywhere