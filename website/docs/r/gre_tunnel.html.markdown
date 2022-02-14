---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_gre_tunnel"
sidebar_current: "docs-cloudflare-resource-gre-tunnel"
description: |-
  Provides a resource which manages GRE tunnels for Magic Transit.
---

# cloudflare_gre_tunnel

Provides a resource, that manages GRE tunnels for Magic Transit.

## Example Usage

```hcl
resource "cloudflare_gre_tunnel" "example" {
  account_id              = "c4a7362d577a6c3019a474fd6f485821"
  name                    = "GRE_1"
  customer_gre_endpoint   = "203.0.113.1"
  cloudflare_gre_endpoint = "203.0.113.1"
  interface_address       = "192.0.2.0/31"
  description             = "Tunnel for ISP X"
  ttl                     = 64
  mtu                     = 1476
  health_check_enabled    = true
  health_check_target     = "203.0.113.1"
  health_check_type       = "reply"
}
```

## Argument Reference

The following arguments are supported:

* `account_id` - (Required) The ID of the account where the tunnel is being created.
* `name` - (Required) Name of the GRE tunnel.
* `customer_gre_endpoint` - (Required) The IP address assigned to the customer side of the GRE tunnel.
* `cloudflare_gre_endpoint` - (Required) The IP address assigned to the Cloudflare side of the GRE tunnel.
* `interface_address` - (Required) 31-bit prefix (/31 in CIDR notation) supporting 2 hosts, one for each side of the tunnel.
* `description` - (Optional) An optional description of the GRE tunnel.
* `ttl` - (Optional) Time To Live (TTL) in number of hops of the GRE tunnel. Minimum value 64. Default: `64`.
* `mtu` - (Optional) Maximum Transmission Unit (MTU) in bytes for the GRE tunnel. Maximum value 1476 and minimum value 576. Default: `1476`.
* `health_check_enabled` - (Optional) Specifies if ICMP tunnel health checks are enabled Default: `true`.
* `health_check_target` - (Optional) The IP address of the customer endpoint that will receive tunnel health checks. Default: `<customer_gre_endpoint>`.
* `health_check_type` - (Optional) Specifies the ICMP echo type for the health check (`request` or `reply`) Default: `reply`.

## Import

An existing GRE tunnel can be imported using the account ID and tunnel ID

```
$ terraform import cloudflare_gre_tunnel.example d41d8cd98f00b204e9800998ecf8427e/cb029e245cfdd66dc8d2e570d5dd3322
```
