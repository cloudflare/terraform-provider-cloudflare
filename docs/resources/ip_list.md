---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_ip_list"
description: Provides IP Lists to be used in Firewall Rules across all zones within the same account.
---

# cloudflare_ip_list

IP Lists are a set of IP addresses or CIDR ranges that are configured on the account level. Once created, IP Lists can be
used in Firewall Rules across all zones within the same account.

## Example Usage

```hcl
resource "cloudflare_ip_list" "example" {
  account_id = "d41d8cd98f00b204e9800998ecf8427e"
  name = "example_list"
  kind = "ip"
  description = "list description"

  item {
    value = "192.0.2.1"
    comment = "Office IP"
  }

  item {
    value = "203.0.113.0/24"
    comment = "Datacenter range"
  }
}
```

## Argument Reference

The following arguments are supported:

- `account_id` - (Required) The ID of the account where the IP List is being created.
- `name` - (Required) The name of the list (used in filter expressions). Valid pattern: `^[a-zA-Z0-9_]+$`. Maximum Length: 50
- `kind` - (Required) The kind of values in the List. Valid values: `ip`.
- `description` - (Optional) A note that can be used to annotate the List. Maximum Length: 500

The **item** block supports:

- `value` - (Required) The IPv4 address, IPv4 CIDR or IPv6 CIDR. IPv6 CIDRs are limited to a maximum of /64.
- `comment` - (Optional) A note that can be used to annotate the item.

## Import

An existing IP List can be imported using the account ID and list ID

```
$ terraform import cloudflare_ip_list.example d41d8cd98f00b204e9800998ecf8427e/cb029e245cfdd66dc8d2e570d5dd3322
```
