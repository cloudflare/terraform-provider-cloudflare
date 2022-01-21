---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_teams_rule"
sidebar_current: "docs-cloudflare-resource-teams-rule"
description: |-
Provides a Cloudflare Teams rule resource.
---

# cloudflare_teams_rule

Provides a Cloudflare Teams rule resource. Teams rules comprise secure web gateway policies.

## Example Usage

```hcl
resource "cloudflare_teams_rule" "rule1" {
  name = "office"
  account_id  = "d57c3de47a013c03ca7e237dd3e61d7d"
  description = "desc"
  precedence = 1
  action = "block"
  filters = ["http"]
  traffic = "http.request.uri == \"https://www.example.com/malicious\""
  rule_settings {
    block_page_enabled = true
    block_page_reason = "access not permitted"
  }
}
```

## Argument Reference

The following arguments are supported:

* `account_id` - (Required) The account to which the teams rule should be added.
* `name` - (Required) The name of the teams rule.
* `description` - (Required) The description of the teams rule.
* `precedence` - (Required) The evaluation precedence of the teams rule.
* `action` - (Required) The action executed by matched teams rule.
* `enabled` - (Optional) Indicator of rule enablement.
* `filters` - (Optional) The protocol or layer to evaluate the traffic and identity expressions.
* `traffic` - (Optional) The wirefilter expression to be used for traffic matching.
* `identity` - (Optional) The wirefilter expression to be used for identity matching.
* `device_posture` - (Optional) The wirefilter expression to be used for device_posture check matching.
* `rule_settings` - (Optional) Additional rule settings.

The **rule_settings** block supports:
* `block_page_enabled` - (Optional) Indicator of block page enablement.
* `block_page_reason` - (Optional) The displayed reason for a user being blocked.
* `override_ips` - (Optional) The IPs to override matching DNS queries with.
* `override_host` - (Optional) The host to override matching DNS queries with.
* `l4override` - (Optional) Settings to forward layer 4 traffic.

The **l4override** block supports:
* `ip` - (Required) Override IP to forward traffic to.
* `port` - (Required) Override Port to forward traffic to.

## Import

Teams Rules can be imported using a composite ID formed of account
ID and teams rule ID.

```
$ terraform import cloudflare_teams_rule.rule1 cb029e245cfdd66dc8d2e570d5dd3322/d41d8cd98f00b204e9800998ecf8427e
```
