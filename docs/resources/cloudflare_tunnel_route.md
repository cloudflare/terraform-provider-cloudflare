---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_tunnel_route"
description: Provides a resource which manages Cloudflare Tunnel Routes for Zero Trust
---

# cloudflare_tunnel_route

Provides a resource, that manages Cloudflare tunnel routes for Zero Trust. Tunnel
routes are used to direct IP traffic through Cloudflare Tunnels.

## Example Usage

```hcl
resource "cloudflare_tunnel_route" "example" {
  account_id = "c4a7362d577a6c3019a474fd6f485821"
  tunnel_id = "f70ff985-a4ef-4643-bbbc-4a0ed4fc8415"
  network = "192.0.2.24/32"
  comment = "New tunnel route for documentation"
}
```

```hcl
resource "cloudflare_argo_tunnel" "tunnel" {
  account_id = "c4a7362d577a6c3019a474fd6f485821"
  name       = "my_tunnel"
  secret     = "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="
}

resource "cloudflare_tunnel_route" "example" {
  account_id = "c4a7362d577a6c3019a474fd6f485821"
  tunnel_id  = cloudflare_argo_tunnel.tunnel.id
  network    = "192.0.2.24/32"
  comment    = "New tunnel route for documentation"
}
```

## Argument Reference

The following arguments are supported:

- `account_id` - (Required) The ID of the account where the tunnel route is being created.
- `tunnel_id` - (Required) The ID of the tunnel that will service the tunnel route.
- `network` - (Required) The IPv4 or IPv6 network that should use this tunnel route, in CIDR notation.
- `comment` - (Optional) Description of the tunnel route.

## Import

An existing tunnel route can be imported using the account ID and network CIDR.

```
$ terraform import cloudflare_tunnel_route c4a7362d577a6c3019a474fd6f485821/192.0.2.24/32
```
