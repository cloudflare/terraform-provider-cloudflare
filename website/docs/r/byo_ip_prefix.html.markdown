---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_byo_ip_prefix"
sidebar_current: "docs-cloudflare-resource-universal-ssl"
description: |-
  Provides the ability to manage Bring-Your-Own-IP prefixes (BYOIP) which are used with or without Magic Transit.
---

# cloudflare_byo_ip_prefix

Provides the ability to manage Bring-Your-Own-IP prefixes (BYOIP) which are used with or without Magic Transit.

## Example Usage

```hcl
resource "cloudflare_byo_ip_prefix" "example" {
    prefix_id = "d41d8cd98f00b204e9800998ecf8427e"
    description = "Example IP Prefix"
    advertisement = "on"
}
```

## Argument Reference

The following arguments are supported:

* `prefix_id` - (Required) The assigned Bring-Your-Own-IP prefix ID.
* `description` - (Optional) The description of the prefix.
* `advertisement` - (Optional) Whether or not the prefix shall be announced. A prefix can be activated or deactivated once every 15 minutes (attempting more regular updates will trigger rate limiting). Valid values: `on` or `off`.

## Import

The current settings for Bring-Your-Own-IP prefixes can be imported using the prefix ID.

```
$ terraform import cloudflare_byo_ip_prefix.example d41d8cd98f00b204e9800998ecf8427e
```
