---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_custom_hostname_fallback_origin"
sidebar_current: "docs-cloudflare_custom_hostname_fallback_origin"
description: !-
  Provides a Cloudflare custom hostname fallback origin resource.
---

# cloudflare_custom_hostname

Provides a Cloudflare custom hostname fallback origin resource.

## Example Usage

```hcl
resource "cloudflare_custom_hostname_fallback_origin" "fallback_origin" {
  zone_id  = "d41d8cd98f00b204e9800998ecf8427e"
  origin   = "fallback.example.com"
}
```

## Argument Reference

The following arguments are supported:

* `zone_id` - (Required) The DNS zone ID where the custom hostname should be assigned.
* `origin` - (Required) Hostname you intend to fallback requests to. Origin must be a proxied A/AAAA/CNAME DNS record within Clouldflare.

## Attributes Reference

The following attribute is exported:

* `status` - Status of the fallback origin's activation.

## Import

Custom hostname fallback origins can be imported using a composite ID formed of the zone ID and [fallback origin](https://api.cloudflare.com/#custom-hostname-fallback-origin-for-a-zone-properties),
separated by a "/" e.g.

```
$ terraform import cloudflare_custom_hostname_fallback_origin.example d41d8cd98f00b204e9800998ecf8427e/fallback.example.com
```
