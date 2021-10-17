---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_argo_tunnel"
sidebar_current: "docs-cloudflare-resource-argo-tunnel"
description: |-
  Provides the ability to manage Cloudflare Argo Tunnels.
---

# cloudflare_argo_tunnel

Argo Tunnel exposes applications running on your local web server on any network with an internet connection without manually adding DNS records or configuring a firewall or router.

## Example Usage

```hcl
resource "cloudflare_argo_tunnel" "example" {
  account_id = "d41d8cd98f00b204e9800998ecf8427e"
  name       = "my-tunnel"
  secret     = "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="
}
```

## Argument Reference

The following arguments are supported:

* `account_id` - (Required) The Cloudflare account ID that you wish to manage the Argo Tunnel on.
* `name` - (Required) A user-friendly name chosen when the tunnel is created. Cannot be empty.
* `secret` - (Required) 32 or more bytes, encoded as a base64 string. The Create Argo Tunnel endpoint sets this as the tunnel's password. Anyone wishing to run the tunnel needs this password.

## Attributes Reference

The following additional attributes are exported:

* `cname` - Usable CNAME for accessing the Argo Tunnel.

## Import

Argo Tunnels can be imported a composite ID of the account ID and tunnel UUID.

-> **Note:** The tunnel secret cannot be imported due to it not being available outside of the creation API calls. It is recommended that you re-create if you don't have the secret saved securely before importing.

```
$ terraform import cloudflare_argo_tunnel.example d41d8cd98f00b204e9800998ecf8427e/fd2455cb-5fcc-4c13-8738-8d8d2605237f
```

where
- `d41d8cd98f00b204e9800998ecf8427e` is the account ID
- `fd2455cb-5fcc-4c13-8738-8d8d2605237f` is the Argo Tunnel UUID
