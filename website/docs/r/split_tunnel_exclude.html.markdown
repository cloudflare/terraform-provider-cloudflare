---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_split_tunnel_exclude"
sidebar_current: "docs-cloudflare-resource-split_tunnel_exclude"
description: |-
  Provides a Cloudflare Split Tunnel Exclude resource.
---

# cloudflare_split_tunnel_exclude

Provides a Cloudflare Split Tunnel Exclude resource. Split tunnel exclude specifies the list of routes excluded from the WARP client's tunnel.

## Example Usage

```hcl
resource "cloudflare_split_tunnel_exclude" "example_split_tunnel_exclude" {
  account_id  = "1d5fdc9e88c8a8c4518b068cd94331fe"
  tunnels {
    host = "*.example.com",
    description = "example domain"
  }
}
```

## Argument Reference

The following arguments are supported:

- `account_id` - (Required) The account to which the device posture rule should be added.
- `tunnels` - (Required) The value of the tunnel attributes. See below for reference
  structure.
- `address` - (Optional) The address in CIDR format to exclude from the tunnel.
- `host` - (Optional) The domain name to exclude from the tunnel.
- `description` - (Optional) The description of the tunnel.

### Tunnel argument

The tunnel structure should have either the address field OR the host field. If both are provided, an error will be returned.
