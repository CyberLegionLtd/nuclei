# Introduction

**Nuclei** is based on the concepts of `YAML` based template files that define how the requests will be sent and processed. This allows easy extensibility capabilities to nuclei. The templates are written in `YAML` which specifies a simple human-readable format to quickly define the execution process for security checks targeting various protocols.

## Getting started with Templates

The two fields that are required in all templates are `id` and `info` fields. Each template begins with declaring these fields and defining appropriate values based on the template being written.

### id

The ID must not container spaces in between. Each ID should also be unique (i.e. it should uniquely specify the purpose of a template). A good ID uniquely identifies what the requests in the template are doing. Let's say you have a template that identifies a git-config file on the webservers, a good name would be `git-config-exposure`. Another example name is `azure-apps-nxdomain-takeover`.

Example of an ID field - 

```yaml
id: git-config
```

```yaml
id: cve-2020-19252
```

```yaml
id: aws-secret-id-exposure
```

### info

The next field is the Information block. It contains `name`, `author`, `severity`, `description`, `reference`, `metadata`, `classification`, `remediation` and `tags`. These fields are described further below - 

- `name` contains the name of the template. 
- `author` contains the authors of the template. 
- `severity` describes the severity of the issue identified by the template. It can be `info`, `low`, `medium`, `high`, `critical` or `unknown`. 
- `description` contains a longer description of the template. `reference` contains optional external links to references about the template. 
- `metadata` can contain a number of arbitrary key-value pairs that provide some useful details about the template. 
- `classification` contains CVE metdata like CVSS, CWE etc in case the template identifies a CVE. 
- `tags` can be used to provide a list of short names that can be used to run the template like `apache,cve` etc.

An example `info` block for a CVE containing many above mentioned fields.

```yaml
info:
  name: VMware Workspace ONE Access - Freemarker SSTI
  author: sherlocksecurity
  severity: critical
  description: An unauthenticated attacker with network access could exploit this vulnerability by sending a specially crafted request to a vulnerable VMware Workspace ONE or Identity Manager. Successful exploitation could result in remote code execution by exploiting a server-side template injection flaw.
  reference:
    - https://www.tenable.com/blog/vmware-patches-multiple-vulnerabilities-in-workspace-one-vmsa-2022-0011
  classification:
    cvss-metrics: CVSS:3.0/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:N/A:N
    cvss-score: 9.8
    cve-id: CVE-2022-22954
    cwe-id: CWE-22
  metadata:
    shodan-query: http.favicon.hash:-1250474341
  tags: cve,cve2022,vmware,ssti,workspaceone
```

Another example `info` block for detecting tugboat configuration file.

```yaml
info:
  name: Tugboat Configuration File Exposure
  description: A Tugboat configuration file was discovered. Tugboat is a command line tool for interacting with DigitalOcean droplets.
  reference:
    - https://github.com/petems/tugboat
    - https://www.digitalocean.com/community/tools/tugboat
  author: geeknik
  severity: critical
  tags: tugboat,config,exposure
```

For more examples, the [nuclei-templates](https://github.com/projectdiscovery/nuclei-templates) template collection can be consulted.