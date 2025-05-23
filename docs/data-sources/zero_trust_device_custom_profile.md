---
page_title: "cloudflare_zero_trust_device_custom_profile Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_zero_trust_device_custom_profile (Data Source)



## Example Usage

```terraform
data "cloudflare_zero_trust_device_custom_profile" "example_zero_trust_device_custom_profile" {
  account_id = "699d98642c564d2e855e9661899b7252"
  policy_id = "f174e90a-fafe-4643-bbbc-4a0ed4fc8415"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String)

### Optional

- `policy_id` (String)

### Read-Only

- `allow_mode_switch` (Boolean) Whether to allow the user to switch WARP between modes.
- `allow_updates` (Boolean) Whether to receive update notifications when a new version of the client is available.
- `allowed_to_leave` (Boolean) Whether to allow devices to leave the organization.
- `auto_connect` (Number) The amount of time in seconds to reconnect after having been disabled.
- `captive_portal` (Number) Turn on the captive portal after the specified amount of time.
- `default` (Boolean) Whether the policy is the default policy for an account.
- `description` (String) A description of the policy.
- `disable_auto_fallback` (Boolean) If the `dns_server` field of a fallback domain is not present, the client will fall back to a best guess of the default/system DNS resolvers unless this policy option is set to `true`.
- `enabled` (Boolean) Whether the policy will be applied to matching devices.
- `exclude` (Attributes List) List of routes excluded in the WARP client's tunnel. (see [below for nested schema](#nestedatt--exclude))
- `exclude_office_ips` (Boolean) Whether to add Microsoft IPs to Split Tunnel exclusions.
- `fallback_domains` (Attributes List) (see [below for nested schema](#nestedatt--fallback_domains))
- `gateway_unique_id` (String)
- `id` (String) The ID of this resource.
- `include` (Attributes List) List of routes included in the WARP client's tunnel. (see [below for nested schema](#nestedatt--include))
- `lan_allow_minutes` (Number) The amount of time in minutes a user is allowed access to their LAN. A value of 0 will allow LAN access until the next WARP reconnection, such as a reboot or a laptop waking from sleep. Note that this field is omitted from the response if null or unset.
- `lan_allow_subnet_size` (Number) The size of the subnet for the local access network. Note that this field is omitted from the response if null or unset.
- `match` (String) The wirefilter expression to match devices. Available values: "identity.email", "identity.groups.id", "identity.groups.name", "identity.groups.email", "identity.service_token_uuid", "identity.saml_attributes", "network", "os.name", "os.version".
- `name` (String) The name of the device settings profile.
- `precedence` (Number) The precedence of the policy. Lower values indicate higher precedence. Policies will be evaluated in ascending order of this field.
- `register_interface_ip_with_dns` (Boolean) Determines if the operating system will register WARP's local interface IP with your on-premises DNS server.
- `sccm_vpn_boundary_support` (Boolean) Determines whether the WARP client indicates to SCCM that it is inside a VPN boundary. (Windows only).
- `service_mode_v2` (Attributes) (see [below for nested schema](#nestedatt--service_mode_v2))
- `support_url` (String) The URL to launch when the Send Feedback button is clicked.
- `switch_locked` (Boolean) Whether to allow the user to turn off the WARP switch and disconnect the client.
- `target_tests` (Attributes List) (see [below for nested schema](#nestedatt--target_tests))
- `tunnel_protocol` (String) Determines which tunnel protocol to use.

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

- `address` (String) The address in CIDR format to include in the tunnel. If `address` is present, `host` must not be present.
- `description` (String) A description of the Split Tunnel item, displayed in the client UI.
- `host` (String) The domain name to include in the tunnel. If `host` is present, `address` must not be present.


<a id="nestedatt--service_mode_v2"></a>
### Nested Schema for `service_mode_v2`

Read-Only:

- `mode` (String) The mode to run the WARP client under.
- `port` (Number) The port number when used with proxy mode.


<a id="nestedatt--target_tests"></a>
### Nested Schema for `target_tests`

Read-Only:

- `id` (String) The id of the DEX test targeting this policy.
- `name` (String) The name of the DEX test targeting this policy.


