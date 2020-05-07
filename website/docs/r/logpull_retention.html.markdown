---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_logpull_retention"
sidebar_current: "docs-cloudflare-resource-logpull-retention"
description: |-
  Allows management of the Logpull Retention settings used to control whether or not to retain HTTP request logs.
---

# cloudflare_logpull_retention

Allows management of the Logpull Retention settings used to control whether or not to retain HTTP request logs.

## Example Usage

```hcl
resource "cloudflare_logpull_retention" "example" {
  zone_id = "fb54f084ca7f7b732d3d3ecbd8ef7bf2"
  enabled = "true"
}
```

## Argument Reference

The following arguments are supported:

* `zone_id` - (Required) The zone ID to apply the log retention to.
* `enabled` - (Required) Whether you wish to retain logs or not.

## Import

You can import existing Logpull Retention using the zone ID as the identifier.

```
$ terraform import cloudflare_logpull_retention.example fb54f084ca7f7b732d3d3ecbd8ef7bf2
```
