---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_devices"
description: Get information on Cloudflare Devices.
---

# cloudflare_devices

Use this data source to lookup [Devices][1].

## Example usage

```hcl
data "cloudflare_devices" "devices" {
  account_id = "c68973221045fe805dfb9aa520153148"
}
```

## Argument Reference

- `account_id` - (Required) The account for which to list the devices.

## Attributes Reference

- `devices` - A list of device object. See below for nested attributes.

**devices**

- `id` - Device ID.
- `key` - The device's public key.
- `device_type` - The type of the device.
- `name` - The device name.
- `version` - The WARP client version.
- `model` - The device model name.
- `os_version` - The operating system version.
- `ip` - IPv4 or IPv6 address.
- `last_seen` - When the device was last seen.
- `created` - When the device was created.
- `updated` - When the device was updated.
- `user_id` - User's ID.
- `user_email` - User's email.
- `user_name` - User's Name.

[1]: https://api.cloudflare.com/#devices-list-devices
