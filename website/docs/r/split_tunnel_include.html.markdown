---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_split_tunnel_include"
sidebar_current: "docs-cloudflare-resource-split_tunnel_include"
description: |-
  Provides a Cloudflare Split Tunnel Include resource.
---

# cloudflare_split_tunnel_include

Provides a Cloudflare Split Tunnel Include resource. Split tunnel include specifies the list of routes allowed in the WARP client's tunnel.

## Example Usage

```hcl
resource "cloudflare_split_tunnel_include" "example_split_tunnel_include" {
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
- `address` - (Optional) The address in CIDR format to include in the tunnel.
- `host` - (Optional) The domain name to include in the tunnel.
- `description` - (Optional) The description of the tunnel.

### Tunnel argument

The tunnel structure should have either the address field OR the host field. If both are provided, an error will be returned.
