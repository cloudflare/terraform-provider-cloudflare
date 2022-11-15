---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_device_policy"
description: Provides a Cloudflare Device Policy resource.
---

# cloudflare_device_policy

Provides a Cloudflare Device Policy resource. Device policies configure settings applied to WARP devices.

## Example Usage

```hcl
resource "cloudflare_device_policy" "developer_warp_policy" {
  account_id = "1d5fdc9e88c8a8c4518b068cd94331fe"
  name = "Developers"
  precedence = 10
  match = "any(identity.groups.name[*] in {\"Developers\"})"
  default = false
  enabled = true
  allow_mode_switch = true
  allow_updates = true
  allowed_to_leave = true
  auto_connect = 0
  captive_portal = 5
  disable_auto_fallback = true
  support_url = "https://cloudflare.com"
  switch_locked = true
  service_mode_v2_mode = "warp"
  service_mode_v2_port = 3000
}
```

## Argument Reference

The following arguments are supported:

- `account_id` - (Required) The account to which the device policy should be added.
- `name` - (Required) Name of the device policy.
- `precedence` - (Optional) The precedence of the device policy. Lower values indicate higher precedence. Policies will be evaluated in ascending order of this field. Cannot be set for the default policy.
- `match` - (Optional) The wirefilter expression to match devices. Cannot be set for the default policy.
- `default` - (Optional) Whether this device policy refers to the default policy.
- `enabled` - (Optional) Whether this device policy is enabled. Cannot be set for the default policy.
- `allow_mode_switch` - (Optional) Whether to allow mode switch for this policy.
- `allow_updates` - (Optional) Whether to allow updates under this policy.
- `allowed_to_leave` - (Optional) Whether to allow devices to leave the organization.
- `auto_connect` - (Optional) The amount of time in minutes to reconnect after having been disabled.
- `captive_portal` - (Optional) The captive portal value for this policy.
- `disable_auto_fallback` - (Optional) Whether to disable auto fallback for this policy.
- `support_url` - (Optional) The support URL that will be opened when sending feedback.
- `switch_locked` - (Optional) Enablement of the ZT client switch lock.
- `service_mode_v2_mode` - (Optional) The service mode.
- `service_mode_v2_port` - (Optional) The port to use for the proxy service mode.

## Attributes Reference

The following additional attributes are exported:

- `id` - ID of the device policy.

## Import

Device policies can be imported using a composite ID formed of account
ID and device policy ID. The default policy does not have an ID but can be
imported with `default` as the policy ID.

```
$ terraform import cloudflare_device_policy.developers cb029e245cfdd66dc8d2e570d5dd3322/0ade592a-62d6-46ab-bac8-01f47c7fa792
$ terraform import cloudflare_device_policy.developers cb029e245cfdd66dc8d2e570d5dd3322/default
```
