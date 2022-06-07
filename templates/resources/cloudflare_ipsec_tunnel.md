---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_ipsec_tunnel"
description: Provides a resource which manages IPsec tunnels for Magic Transit.
---

# cloudflare_ipsec_tunnel

Provides a resource, that manages IPsec tunnels for Magic Transit.

## Example Usage

```hcl
resource "cloudflare_ipsec_tunnel" "example" {
  account_id          = "c4a7362d577a6c3019a474fd6f485821"
  name                = "IPsec_1"
  customer_endpoint   = "203.0.113.1"
  cloudflare_endpoint = "203.0.113.1"
  interface_address   = "192.0.2.0/31"
  description         = "Tunnel for ISP X"
}
```

## Argument Reference

The following arguments are supported:

- `account_id` - (Required) The ID of the account where the tunnel is being created.
- `name` - (Required) Name of the IPsec tunnel.
- `customer_endpoint` - (Required) IP address assigned to the customer side of the IPsec tunnel.
- `cloudflare_endpoint` - (Required) IP address assigned to the Cloudflare side of the IPsec tunnel.
- `interface_address` - (Required) 31-bit prefix (/31 in CIDR notation) supporting 2 hosts, one for each side of the tunnel.
- `description` - (Optional) An optional description of the IPsec tunnel.

## Import

An existing IPsec tunnel can be imported using the account ID and tunnel ID

```
$ terraform import cloudflare_ipsec_tunnel.example d41d8cd98f00b204e9800998ecf8427e/cb029e245cfdd66dc8d2e570d5dd3322
```
