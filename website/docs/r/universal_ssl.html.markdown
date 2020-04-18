---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_universal_ssl"
sidebar_current: "docs-cloudflare-resource-universal-ssl"
description: |-
  Provides the ability to manage Universal SSL status.
---

# cloudflare_universal_ssl

Provides the ability to enable/disable Universal SSL for a given zone.

## Example Usage

```hcl
resource "cloudflare_universal_ssl" "example" {
    zone_id = "d41d8cd98f00b204e9800998ecf8427e"
    settings {
        status = "on"
    }
}
```

## Argument Reference

The following arguments are supported:

* `zone_id` - (Required) The DNS zone ID to apply to.
* `status` - (Required) The Universal SSL status. ["on", "off"]


## Import

The current Universal SSL status can be imported using the zone's ID.

```
$ terraform import cloudflare_universal_ssl.example d41d8cd98f00b204e9800998ecf8427e
```
