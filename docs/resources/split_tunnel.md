---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_split_tunnel"
description: Provides a Cloudflare Split Tunnel resource.
---

# cloudflare_split_tunnel

Provides a Cloudflare Split Tunnel resource. Split tunnels are used to either
include or exclude lists of routes from the WARP client's tunnel.

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

# Create a device policy
resource "cloudflare_device_policy" "developer_warp_policy" {
  account_id  = "1d5fdc9e88c8a8c4518b068cd94331fe"
  name        = "Developers"
  precedence = 10
  match = "any(identity.groups.name[*] in {\"Developers\"})"
  switch_locked = true
}

# Excluding *.example.com from WARP routes for a particular device policy
resource "cloudflare_split_tunnel" "example_device_policy_split_tunnel_exclude" {
  account_id = "1d5fdc9e88c8a8c4518b068cd94331fe"
  policy_id  = cloudflare_device_policy.developer_warp_policy.id
  mode       = "exclude"
  tunnels {
    host        = "*.example.com",
    description = "example domain"
  }
}

# Including *.example.com in WARP routes for a particular device policy
resource "cloudflare_split_tunnel" "example_split_tunnel_include" {
  account_id = "1d5fdc9e88c8a8c4518b068cd94331fe"
  policy_id  = cloudflare_device_policy.developer_warp_policy.id
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
- `mode` - (Required) The split tunnel mode. Valid values are `include` or `exclude`.
- `tunnels` - (Required) The value of the tunnel attributes (refer to the [nested schema](#nestedblock--tunnels)).
- `policy_id` - (Optional) The device policy ID with which to associate this split tunnel configuration. If missing, will refer to the default device policy.

<a id="nestedblock--tunnels"></a>
**Nested schema for `tunnels`**

- `address` - (Optional) The address in CIDR format to include in the tunnel configuration. Conflicts with `"host"`.
- `host` - (Optional) The domain name to include in the tunnel configuration. Conflicts with `"address"`.
- `description` - (Optional) The description of the tunnel.

## Import

Split Tunnels can be imported using the account identifer, policy ID, and mode. Split Tunnels for default device policies must use "default" as the policy ID.

```
$ terraform import cloudflare_split_tunnel.example 1d5fdc9e88c8a8c4518b068cd94331fe/default/exclude
$ terraform import cloudflare_split_tunnel.example 1d5fdc9e88c8a8c4518b068cd94331fe/0ade592a-62d6-46ab-bac8-01f47c7fa792/exclude
```
