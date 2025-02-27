---
page_title: "cloudflare_zero_trust_device_default_profile Resource - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_zero_trust_device_default_profile (Resource)



## Example Usage

```terraform
resource "cloudflare_zero_trust_device_default_profile" "example_zero_trust_device_default_profile" {
  account_id = "699d98642c564d2e855e9661899b7252"
  allow_mode_switch = true
  allow_updates = true
  allowed_to_leave = true
  auto_connect = 0
  captive_portal = 180
  disable_auto_fallback = true
  exclude_office_ips = true
  service_mode_v2 = {
    mode = "proxy"
    port = 3000
  }
  support_url = "https://1.1.1.1/help"
  switch_locked = true
  tunnel_protocol = "wireguard"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String)

### Optional

- `allow_mode_switch` (Boolean) Whether to allow the user to switch WARP between modes.
- `allow_updates` (Boolean) Whether to receive update notifications when a new version of the client is available.
- `allowed_to_leave` (Boolean) Whether to allow devices to leave the organization.
- `auto_connect` (Number) The amount of time in seconds to reconnect after having been disabled.
- `captive_portal` (Number) Turn on the captive portal after the specified amount of time.
- `disable_auto_fallback` (Boolean) If the `dns_server` field of a fallback domain is not present, the client will fall back to a best guess of the default/system DNS resolvers unless this policy option is set to `true`.
- `exclude_office_ips` (Boolean) Whether to add Microsoft IPs to Split Tunnel exclusions.
- `service_mode_v2` (Attributes) (see [below for nested schema](#nestedatt--service_mode_v2))
- `support_url` (String) The URL to launch when the Send Feedback button is clicked.
- `switch_locked` (Boolean) Whether to allow the user to turn off the WARP switch and disconnect the client.
- `tunnel_protocol` (String) Determines which tunnel protocol to use.

### Read-Only

- `default` (Boolean) Whether the policy will be applied to matching devices.
- `enabled` (Boolean) Whether the policy will be applied to matching devices.
- `exclude` (Attributes List) (see [below for nested schema](#nestedatt--exclude))
- `fallback_domains` (Attributes List) (see [below for nested schema](#nestedatt--fallback_domains))
- `gateway_unique_id` (String)
- `id` (String) The ID of this resource.
- `include` (Attributes List) (see [below for nested schema](#nestedatt--include))

<a id="nestedatt--service_mode_v2"></a>
### Nested Schema for `service_mode_v2`

Optional:

- `mode` (String) The mode to run the WARP client under.
- `port` (Number) The port number when used with proxy mode.


<a id="nestedatt--exclude"></a>
### Nested Schema for `exclude`

Read-Only:

- `address` (String) The address in CIDR format to exclude from the tunnel. If `address` is present, `host` must not be present.
- `description` (String) A description of the Split Tunnel item, displayed in the client UI.
- `host` (String) The domain name to exclude from the tunnel. If `host` is present, `address` must not be present.


<a id="nestedatt--fallback_domains"></a>
### Nested Schema for `fallback_domains`

Read-Only:

- `description` (String) A description of the fallback domain, displayed in the client UI.
- `dns_server` (List of String) A list of IP addresses to handle domain resolution.
- `suffix` (String) The domain suffix to match when resolving locally.


<a id="nestedatt--include"></a>
### Nested Schema for `include`

Read-Only:

- `address` (String) The address in CIDR format to include in the tunnel. If address is present, host must not be present.
- `description` (String) A description of the split tunnel item, displayed in the client UI.
- `host` (String) The domain name to include in the tunnel. If host is present, address must not be present.

## Import

Import is supported using the following syntax:

```shell
$ terraform import cloudflare_zero_trust_device_default_profile.example '<account_id>'
```
