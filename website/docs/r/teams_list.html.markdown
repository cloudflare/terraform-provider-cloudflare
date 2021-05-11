---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_teams_list"
sidebar_current: "docs-cloudflare-resource-teams-list"
description: |-
  Provides a Cloudflare Teams List resource.
---

# cloudflare_teams_list

Provides a Cloudflare Teams List resource. Teams lists are referenced when creating secure web gateway policies or device posture rules.

## Example Usage

```hcl
resource "cloudflare_teams_list" "corporate_devices" {
  account_id  = "1d5fdc9e88c8a8c4518b068cd94331fe"
  name        = "Corporate devices"
  type        = "SERIAL"
  description = "Serial numbers for all corporate devices."
  items       = ["8GE8721REF", "5RE8543EGG", "1YE2880LNP"]
}
```

## Argument Reference

The following arguments are supported:

* `account_id` - (Required) The account to which the teams list should be added.
* `name` - (Required) Name of the teams list.
* `type` - (Required) The teams list type. Valid values are `SERIAL`, `URL`, `DOMAIN`, and `EMAIL`.
* `items` - (Required) The items of the teams list.
* `description` - (Optional) The description of the teams list.

## Attributes Reference

The following additional attributes are exported:

* `id` - ID of the teams list.

## Import

Teams lists can be imported using a composite ID formed of account
ID and teams list ID.

```
$ terraform import cloudflare_teams_list.corporate_devices cb029e245cfdd66dc8d2e570d5dd3322/d41d8cd98f00b204e9800998ecf8427e
```
