---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_split_tunnel"
sidebar_current: "docs-cloudflare-resource-split-tunnel"
description: |-
  Provides a Cloudflare Split Tunnel resource.
---

# cloudflare_split_tunnel

Provides a Cloudflare Split Tunnel resource. Split tunnels are used to either
include or exclude lists of routes from the WARP client's tunnel.  A
single resource should be used choosing either the `include` or `exclude` mode.

## Example Usage

```hcl
# Excluding *.example.com from WARP routes
resource "cloudflare_split_tunnel" "example_split_tunnel_exclude" {
  account_id = "1d5fdc9e88c8a8c4518b068cd94331fe"
  mode       = "exclude"
  tunnels {
    host        = "*.example.com",
    description = "example domain"
  }
}

# Including *.example.com in WARP routes
resource "cloudflare_split_tunnel" "example_split_tunnel_include" {
  account_id = "1d5fdc9e88c8a8c4518b068cd94331fe"
  mode       = "include"
  tunnels {
    host        = "*.example.com",
    description = "example domain"
  }
}
```

## Argument Reference

The following arguments are supported:

- `account_id` - (Required) The account to which the device posture rule should be added.
- `mode` - (Required) The split tunnel mode.  Valid values are `include` or `exclude`.
- `tunnels` - (Required) The value of the tunnel attributes (refer to the [nested schema](#nestedblock--tunnels)).

<a id="nestedblock--tunnels"></a>
**Nested schema for `tunnels`**

- `address` - (Optional) The address in CIDR format to include in the tunnel configuration. Conflicts with `"host"`.
- `host` - (Optional) The domain name to include in the tunnel configuration. Conflicts with `"address"`.
- `description` - (Optional) The description of the tunnel.

## Import

Split Tunnels can be imported using the account identifer and mode.

```
$ terraform import cloudflare_split_tunnel.example 1d5fdc9e88c8a8c4518b068cd94331fe/exclude
```
