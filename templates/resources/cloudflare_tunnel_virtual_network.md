---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_tunnel_virtual_network"
description: Provides a resource which manages Cloudflare Tunnel Virtual Networks for Zero Trust
---

# cloudflare_tunnel_virtual_network

Provides a resource, that manages Cloudflare tunnel virtual networks for Zero Trust. Tunnel
virtual networks are used for segregation of Tunnel IP Routes via Virtualized Networks to 
handle overlapping private IPs in your origins..

## Example Usage

```hcl
resource "cloudflare_tunnel_virtual_network" "example" {
  account_id = "c4a7362d577a6c3019a474fd6f485821"
  name = "vnet-for-documentation"
  comment = "New tunnel virtual network for documentation"
}
```

## Argument Reference

The following arguments are supported:

- `account_id` - (Required) The ID of the account where the tunnel virtual network is being created.
- `name` - (Required) A user-friendly name chosen when the virtual network is created.
- `is_default_network` - (Optional) Whether this virtual network is the default one for the account. This means IP Routes belong to this virtual network and Teams Clients in the account route through this virtual network, unless specified otherwise for each case.
- `comment` - (Optional) Description of the tunnel virtual network.

## Import

An existing tunnel virtual networks can be imported using the account ID and virtual network ID.

```
$ terraform import cloudflare_tunnel_virtual_network c4a7362d577a6c3019a474fd6f485821/3c8ff8af-b487-45bd-89e3-4c85a1532600
```
