---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_argo"
sidebar_current: "docs-cloudflare-resource-argo"
description: |-
  Provides the ability to manage Cloudflare Argo features.
---

# cloudflare_argo

Cloudflare Argo controls the routing to your origin and tiered caching options to speed up your website browsing experience.

## Example Usage

```hcl
resource "cloudflare_argo" "example" {
  zone_id        = "d41d8cd98f00b204e9800998ecf8427e"
  tiered_caching = "on"
  smart_routing  = "on"
}
```

## Argument Reference

The following arguments are supported:

* `zone_id` - (Required) The DNS zone ID that you wish to manage Argo on.
* `tiered_caching` - (Optional) Whether tiered caching is enabled. Valid values: `on` or `off`. Defaults to `off`.
* `smart_routing` - (Optional) Whether smart routing is enabled. Valid values: `on` or `off`. Defaults to `off`.


## Import

Argo settings can be imported the zone ID.

```
$ terraform import cloudflare_argo.example d41d8cd98f00b204e9800998ecf8427e
```

where `d41d8cd98f00b204e9800998ecf8427e` is the zone ID.
