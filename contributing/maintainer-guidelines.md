# Maintainer guidelines

## Labels

`kind/bug` and `kind/enhancement` are based on the type of issue opened.
`needs-triage` is automatically added to all new issues.

| Label | Purpose |
|-------|---------|
| `kind/breaking-change` | |
| `kind/bug` | Categorizes issue or PR as related to a bug. |
| `kind/crash` | Categorizes issue or PR as related to a crash caused by the provider. |
| `kind/documentation` | Categorizes issue or PR as related to documentation. |
| `kind/enhancement` | Categorizes issue or PR as related to improving an existing feature. |
| `kind/failing-test` | Categorizes issue or PR as related to a consistently or frequently failing test. |
| `kind/flakey` | Categorizes issue or PR as related to a flaky test. |
| `kind/new-data-source` | Categorizes issue or PR as related to needing to create a datasource. |
| `kind/new-resource` | Categorizes issue or PR as related to creating a new resource. |
| `kind/question` | Categorizes issue or PR as related to a question about the provider or the use of the provider. |
| `kind/regression` | Categorizes issue or PR as related to a regression from a prior release. |
| `kind/support` | Categorizes issue or PR as related to user support. |
| `lifecycle/stale` | |
| `likelihood/all` | Categorizes issue or PR as impacting all users. |
| `likelihood/few` | Categorizes issue or PR as impacting a small portion of users. |
| `likelihood/low` | Categorizes issue or PR as impacting a low portion of users. |
| `likelihood/many` | Categorizes issue or PR as impacting many users. |
| `likelihood/most` | Categorizes issue or PR as impacting most users. |
| `needs-triage` | Indicates an issue or PR lacks a `triage/foo` label and requires one. |
| `service/access` | Categorizes issue or PR as related to the Access service. |
| `service/argo` | Categorizes issue or PR as related to the Argo service. |
| `service/custom_pages` | Categorizes issue or PR as related to the custom pages service. |
| `service/dns` | Categorizes issue or PR as related to the DNS service. |
| `service/firewall` | Categorizes issue or PR as related to the Firewall service. |
| `service/iam` | Categorizes issue or PR as related to the IAM service. |
| `service/load_balancing` | Categorizes issue or PR as related to the Load Balancing service. |
| `service/magic_transit` | Categorizes issue or PR as related to Magic Transit suite of services. |
| `service/notifications` | Categorizes issue or PR as related to the notification service. |
| `service/page_rules` | Categorizes issue or PR as related to the Page Rules service. |
| `service/rulesets` | Categorizes issue or PR as related to the Rulesets service. |
| `service/spectrum` | Categorizes issue or PR as related to the Spectrum service. |
| `service/ssl` | Categorizes issue or PR as related to the SSL service. |
| `service/tunnel` | Categorizes issue or PR as related to the Tunnel service. |
| `service/workers` | Categorizes issue or PR as related to the Workers service. |
| `service/zones` | Categorizes issue or PR as related to the Zones service. |
| `triage/accepted` | Indicates an issue or PR is ready to be actively worked on. |
| `triage/duplicate` | Indicates an issue is a duplicate of other open issue. |
| `triage/needs-information` | Indicates an issue needs more information in order to work on it. |
| `triage/not-reproducible` | Indicates an issue can not be reproduced as described. |
| `triage/unresolved` | Indicates an issue that can not or will not be resolved. |
| `workflow/needs-review` | Indicates an issue or PR needs review or feedback. |
| `workflow/pending-cloudflare-response` | Indicates an issue or PR requires a response from the Cloudflare team. |
| `workflow/pending-contributor-response` | Indicates an issue or PR requires a response from a contributor. |
| `workflow/pending-maintainer-response` | Indicates an issue or PR requires a response from the maintainer team. |
| `workflow/pending-op-response` | Indicates an issue or PR requires a response from the original poster. |
| `workflow/pending-public-documentation `| Indicates an issue or PR requires changes to public documentation confirming suitability for use. |
| `workflow/pending-upstream-library` | Indicates an issue or PR requires changes from an upstream library. |
| `workflow/pr-attached` | Indicates the issue has PR(s) attached. |
| `workflow/skip-changelog-entry` | Denotes an issue or PR does not require a changelog entry to be added before merging. |


## Tasks

### Regularly

- Triage issues and label them aiming to include at least on of each ("kind",
  "triage" and "workflow") to communicate the state of the issue.
- Review open PRs
  - Running acceptance tests for PRs locally.
  - Marking the PR as "approved" or "request changes" with comments on the
    changes to be made.
  - Scan open issues to see if any can be linked to others for better visibility.
  - Ensure any changes (such as `cloudflare-go`) are co-ordinated for merging.
  - Should the original creator be unresponsive, determine if the PR priority is
    worth finishing it yourself.
  - Confirm PR has [CHANGELOG entry](changelog-process.md) where it makes sense.
- Follow up on open PRs. Stale issues will automatically close thanks to
  automation.

### Fortnightly

- [Cut a release](release-process.md) (including `cloudflare-go` dependencies).
  Releases can be more frequent if needed however acceptance testing must be
  carried out for each release.

## Questions?

Hit up @jacobbednarz
