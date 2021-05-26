---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_device_posture_rule"
sidebar_current: "docs-cloudflare-resource-device-posture-rule"
description: |-
  Provides a Cloudflare Device Posture Rule resource.
---

# cloudflare_device_posture_rule

Provides a Cloudflare Device Posture Rule resource. Device posture rules configure security policies for device posture checks.

## Example Usage

```hcl
resource "cloudflare_device_posture_rule" "corporate_devices_posture_rule" {
  account_id  = "1d5fdc9e88c8a8c4518b068cd94331fe"
  name        = "Corporate devices posture rule"
  type        = "serial_number"
  description = "Device posture rule for corporate devices."
  schedule    = "24h"
  match {
    platform = "mac"
  }
  input {
    id = cloudflare_teams_list.corporate_devices.id
  }
}
```

## Argument Reference

The following arguments are supported:

* `account_id` - (Required) The account to which the device posture rule should be added.
* `type` - (Required) The device posture rule type. Valid values are `file`, `application`, and `serial_number`.
* `input` - (Required) The value to be checked against. See below for reference
  structure.
* `name` - (Optional) Name of the device posture rule.
* `schedule` - (Optional) Tells the client when to run the device posture check.
  Must be in the format `"1h"` or `"30m"`. Valid units are `h` and `m`.
* `description` - (Optional) The description of the device posture rule.
* `match` - (Optional) The conditions that the client must match to run the rule. See below for reference structure.

### Match argument

The match structure allows the following:

* `platform` - (Required) The platform of the device. Valid values are `windows`, `mac`, `linux`, `android`, and `ios`.

### Input argument

The input structure depends on the device posture rule type.

**serial_number** allows the following:

* `id` - (Required) The Teams List id.

**file** allows the following:

* `path` - (Required) The path to the file.
* `exists` - (Optional) Checks if the file should exist.
* `thumbprint` - (Optional) The thumbprint of the file certificate.
* `sha256` - (Optional) The sha256 hash of the file.

**application** allows the following:

* `path` - (Required) The path to the application.
* `thumbprint` - (Optional) The thumbprint of the application certificate.
* `running` - (Optional) Checks if the application should be running.

## Attributes Reference

The following additional attributes are exported:

* `id` - ID of the device posture rule.

## Import

Device posture rules can be imported using a composite ID formed of account
ID and device posture rule ID.

```
$ terraform import cloudflare_device_posture_rule.corporate_devices cb029e245cfdd66dc8d2e570d5dd3322/d41d8cd98f00b204e9800998ecf8427e
```
