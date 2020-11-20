---
name: "\U0001F41B Bug report"
about: When something isn't working as expected or documented
labels: kind/bug, needs-triage
---

<!--
Hi there,

Thank you for opening an issue. Please note that we try to keep the Terraform
issue tracker reserved for bug reports and feature requests. For general usage
questions, please see: https://www.terraform.io/community.html.
-->

### Terraform version

<!--
Run `terraform -v` to show the version of Terraform and providers you are using.
If you are not running the latest version of Terraform or the providers, please
upgrade because your issue may have already been fixed.
-->

### Affected resource(s)

<!--
Please list the resources as a list, for example:
- opc_instance
- opc_storage_volume

If this issue appears to affect multiple resources, it may be an issue with
Terraform's core, so please mention this.
-->

### Terraform configuration files

<!--
```hcl
# Copy-paste your Terraform configurations here - for large Terraform configs,
# please use a service like Dropbox and share a link to the ZIP file. For
# security, you can also encrypt the files using our GPG public key.
```
-->

### Debug output

<!--
Please provider a link to a GitHub Gist containing the complete debug output:
https://www.terraform.io/docs/internals/debugging.html. Please do NOT paste the
debug output in the issue; just paste a link to the Gist. To speed up the process
please provide the *entire* debug log instead of a snippet. You should redact any 
sensitive details before posting it.
-->

### Panic output

<!--
If Terraform produced a panic, please provide a link to a GitHub Gist containing
the output of the `crash.log`.
-->

### Expected behavior

<!--
What should have happened?
-->

### Actual behavior

<!--
What actually happened?
-->

### Steps to reproduce

<!--
Please list the steps required to reproduce the issue, for example:
1. `terraform apply`
-->

### Important factoids

<!--
Are there anything atypical about your accounts that we should know? For
example: Running in EC2 Classic? Custom version of OpenStack? Tight ACLs?
-->

### References

<!--
Are there any other GitHub issues (open or closed) or Pull Requests that should
be linked here? For example:

- GH-1234
-->

### Community note

* Please vote on this issue by adding a 👍 [reaction](https://blog.github.com/2016-03-10-add-reactions-to-pull-requests-issues-and-comments/)
  to the original issue to help the community and maintainers prioritize this request
* If you are interested in working on this issue or have submitted a pull
  request, please leave a comment

<!--- Thank you for keeping this note for the community --->
