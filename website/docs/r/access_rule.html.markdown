---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_access_rule"
sidebar_current: "docs-cloudflare-resource-access-rule"
description: |-
  Provides a Cloudflare IP Firewall Access Rule resource.
---

# cloudflare_access_rule

Provides a Cloudflare IP Firewall Access Rule resource. Access control can be applied on basis of IP addresses, IP ranges, AS numbers or countries.

## Example Usage

```hcl
# Challenge requests coming from known Tor exit nodes.
resource "cloudflare_access_rule" "tor_exit_nodes" {
  notes = "Requests coming from known Tor exit nodes"
  mode = "challenge"
  configuration = {
    target = "country"
    value = "T1"
  }
}

# Whitelist (sic!) requests coming from Antarctica, but only for single zone.
resource "cloudflare_access_rule" "antarctica" {
  notes = "Requests coming from Antarctica"
  mode = "whitelist"
  configuration = {
    target = "country"
    value = "AQ"
  }
  zone_id = "cb029e245cfdd66dc8d2e570d5dd3322"
}

# Whitelist office's network IP ranges on all account zones (or other lists of resources).
# Resulting Terraform state will be a list of resources.
provider "cloudflare" {
  # ... other provider configuration
  account_id = "d41d8cd98f00b204e9800998ecf8427e"
}
variable "my_office" {
  type = "list"
  default = ["192.0.2.0/24", "198.51.100.0/24", "2001:db8::/56"]
}
resource "cloudflare_access_rule" "office_network" {
  count = length(var.my_office)
  notes = "Requests coming from office network"
  mode = "whitelist"
  configuration = {
    target = "ip_range"
    value = element(var.my_office, count.index)
  }
}
```

## Argument Reference

The following arguments are supported:

* `zone_id` - (Optional) The DNS zone to which the access rule should be added.
* `mode` - (Required) The action to apply to a matched request. Allowed values: "block", "challenge", "whitelist", "js_challenge"
* `notes` - (Optional) A personal note about the rule. Typically used as a reminder or explanation for the rule.
* `configuration` - (Required) Rule configuration to apply to a matched request. It's a complex value. See description below.

**Note:** If both `zone` and `zone_id` are empty, then access rule will be set to the account level and apply to all their zones.

The **configuration** block supports:

* `target` - (Required) The request property to target. Allowed values: "ip", "ip6", "ip_range", "asn", "country"
* `value` - (Required) The value to target. Depends on target's type.

## Attributes Reference

The following attributes are exported:

* `id` - The access rule ID.
* `zone_id` - The DNS zone ID.

## Import

Records can be imported using a composite ID formed of access rule type,
access rule type identifier and identifer value, e.g.

```
$ terraform import cloudflare_access_rule.default zone/cb029e245cfdd66dc8d2e570d5dd3322/d41d8cd98f00b204e9800998ecf8427e
```

where:

* `zone` - access rule type (`account`, `zone` or `user`)
* `cb029e245cfdd66dc8d2e570d5dd3322` - access rule type ID (i.e the zone ID
  or account ID you wish to target)
* `d41d8cd98f00b204e9800998ecf8427e` - access rule ID as returned by
  respective API endpoint for the type you are attempting to import.
