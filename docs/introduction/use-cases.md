# Use Cases

Nuclei is a flexible scanning tool that can be used by Bug Bounty Hunters, Penetration Testers, Blue Teamers, etc for acheiving their scanning requirements.

## For Bug Bounty Hunters

Nuclei allows you to customise your testing approach with your own suite of checks and easily run across your bug bounty programs. Moreover, Nuclei can be easily integrated into any continuous scanning workflow.

- Designed to be easily integrated into other tool workflow.
- Can process thousands of hosts in few minutes.
- Easily automate your custom testing approach with our simple YAML DSL.

Please check our other open-source projects that might fit into your bug bounty workflow: [github.com/projectdiscovery](https://github.com/projectdiscovery), we also host daily refresh of DNS data at [Chaos](http://chaos.projectdiscovery.io/).

## For Penetration Testers

Nuclei immensely improve how you approach security assessment by augmenting the manual, repetitive processes. Consultancies are already converting their manual assessment steps with Nuclei, it allows them to run set of their custom assessment approach across thousands of hosts in an automated manner.

Pen-testers get the full power of our public templates and customization capabilities to speed up their assessment process, and specifically with the regression cycle where you can easily verify the fix.

- Easily create your compliance, standards suite (e.g. OWASP Top 10) checklist.
- With capabilities like [fuzz](https://nuclei.projectdiscovery.io/templating-guide/#advance-fuzzing) and [workflows](https://nuclei.projectdiscovery.io/templating-guide/#workflows), complex manual steps and repetitive assessment can be easily automated with Nuclei.
- Easy to re-test vulnerability-fix by just re-running the template.

## For Developers and Organisations

Nuclei is built with simplicity in mind, with the community backed templates by hundreds of security researchers, it allows you to stay updated with the latest security threats using continuous Nuclei scanning on the hosts. It is designed to be easily integrated into regression tests cycle, to verify the fixes and eliminate vulnerabilities from occurring in the future.

- CI/CD: Engineers are already utilising Nuclei within their CI/CD pipeline, it allows them to constantly monitor their staging and production environments with customised templates.
- Continuous Regression Cycle: With Nuclei, you can create your custom template on every new identified vulnerability and put into Nuclei engine to eliminate in the continuous regression cycle.
  
We have [a discussion thread around this](https://github.com/projectdiscovery/nuclei#:~:text=a%20discussion%20thread%20around%20this), there are already some bug bounty programs giving incentives to hackers on writing nuclei templates with every submission, that helps them to eliminate the vulnerability across all their assets, as well as to eliminate future risk in reappearing on productions. If you're interested in implementing it in your organisation, feel free to [reach out to us](contact@projectdiscovery.io). We will be more than happy to help you in the getting started process, or you can also post into the [discussion thread for any help](https://github.com/projectdiscovery/nuclei-templates/discussions/693).

![workflow](../images/regression-with-nuclei.jpeg)
