---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_magic_firewall_ruleset"
description: Provides the ability to manage a Magic Firewall Ruleset and it's firewall rules which are used with Magic Transit.
---

# cloudflare_magic_firewall_ruleset

Magic Firewall is a network-level firewall to protect networks that are onboarded to Cloudflare's Magic Transit. This resource
creates a root ruleset on the account level and contains one or more rules. Rules can be crafted in Wireshark syntax and
are evaluated in order, with the first rule having the highest priority.

## Example Usage

```hcl
resource "cloudflare_magic_firewall_ruleset" "example" {
  account_id = "d41d8cd98f00b204e9800998ecf8427e"
  name = "Magic Transit Ruleset"
  description = "Global mitigations"

  rules = [
    {
      action = "allow"
      expression = "tcp.dstport in { 32768..65535 }"
      description = "Allow TCP Ephemeral Ports"
      enabled = "true"
    },
    {
      action = "block"
      expression = "ip.len >= 0"
      description = "Block all"
      enabled = "true"
    }
  ]
}
```

## Argument Reference

The following arguments are supported:

- `account_id` - (Required) The ID of the account where the ruleset is being created.
- `name` - (Required) The name of the ruleset.
- `description` - (Optional) A note that can be used to annotate the ruleset.

The **rules** block is a list of maps with the following attributes:

- `action` - (Required) Valid values: `allow` or `block`.
- `expression` - (Required) A Firewall expression using Wireshark syntax.
- `description` - (Optional) A note that can be used to annotate the rule.
- `enabled` - (Required) Whether the rule is enabled or not. Valid values: `true` or `false`.

## Import

An existing Magic Firewall Ruleset can be imported using the account ID and ruleset ID

```
$ terraform import cloudflare_magic_firewall_ruleset.example d41d8cd98f00b204e9800998ecf8427e/cb029e245cfdd66dc8d2e570d5dd3322
```
