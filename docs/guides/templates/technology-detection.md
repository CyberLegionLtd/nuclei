# Technology Detection Templates

Nuclei makes it very easy to write templates to fingerprint the applications running on the target. Technology specific templates can be created to detect an applications, which can also allow you to go as far as to detect the versions installed too. Just find a fixed string present on the target application of your choice and/or a path that’s only present in that application and create a template.

For example, we know that Jira installations contain a path called `/secure/Dashboard.jspa` and the response to that path contains a fixed string - `Project Management Software`. Utilising these two pieces of information, we can write a template to detect running Jira servers as follows -

```yaml
id: jira-detect

info:
  name: Detect Jira Issue Management Software
  author: bauthard
  severity: informative

requests:
  - method: GET
    path:
      - "{{BaseURL}}/secure/Dashboard.jspa"
      - "{{BaseURL}}/jira/secure/Dashboard.jspa"
    matchers:
      - type: word
        part: body
        words:
          - "Project Management Software"
```

Similarly, Jenkins Server has a header present in the response called `x-jenkins`. Utilising this information, we can also write a template to detect Jenkins servers very easily by searching the header part for the word.

```yaml
id: jenkins-headers-detect

info:
  name: Jenkins Headers Based Detection
  author: ice3man
  severity: informative

requests:
  - method: GET
    path:
      - "{{BaseURL}}/"
    matchers:
      - type: word
        words:
          - "X-Jenkins"
        part: header
```

The same technique can be applied to create templates to detect any running piece of application you desire. All you need is a unique pattern present somewhere in the application, and you can detect it effortlessly.