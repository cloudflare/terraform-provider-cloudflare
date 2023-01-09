---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_device_managed_networks"
description: Provides a Cloudflare Device Managed Networks resource.
---

# cloudflare_device_managed_networks

Provides a Cloudflare Device Managed Networks resource. Device managed networks allow for building location-aware device settings policies.

## Example Usage

```hcl
resource "cloudflare_device_managed_networks" "managed_networks" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  name        = "managed-network-1"
  type        = "tls"
  config {
      tls_sockaddr       = "foobar:1234"
      sha256 = "b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c"
  }
}

```

## Argument Reference

The following arguments are supported:

- `account_id` - (Required) The account to which the device managed network should be added.
- `name` - (Required)The name of the Device Managed Network. Must be unique.
- `type` - (Required) The type of Device Managed Network. Valid values is `tls`.
- `config` - (Required) The configuration object containing information for the WARP client to detect the managed network.

### Config argument

- `tls_sockaddr` - (Required) The third-party API's URL.
- `sha256` - (Required) The third-party authorization API URL.

## Import

Import is supported using the following syntax:

```shell
$ terraform import cloudflare_device_managed_networks.example <account_id>/<device_managed_networks_id>
```