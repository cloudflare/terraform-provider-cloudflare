---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_worker_script"
sidebar_current: "docs-cloudflare-resource-worker-script"
description: |-
  Provides a Cloudflare worker script resource.
---

# cloudflare_worker_script

Provides a Cloudflare worker script resource. In order for a script to be active, you'll also need to setup a `cloudflare_worker_route`.

## Example Usage

__NOTE:__ This is for non-enterprise accounts where there is one script per zone. For enterprise accounts, see the "multi-script" example below.

```hcl
# Sets the script for the example.com zone
resource "cloudflare_worker_script" "my_script" {
  zone_id = "ae36f999674d196762efcc5abb06b345"
  content = "${file("script.js")}"
}
```

## Multi-script example usage

__NOTE:__ This is only for enterprise accounts. With multi-script, each script is given a `name` instead of a `zone_id`

```hcl
# Sets the script with the name "script_1"
resource "cloudflare_worker_script" "my_script" {
  name = "script_1"
  content = "${file("script.js")}"
}
```

## Argument Reference

The following arguments are supported:

* `zone_id` - (Required for single-script accounts) The zone ID for the script.
* `name` - (Required for multi-script accounts) The name for the script.
* `content` - (Required) The script content.

## Import

### single-script

To import a script from a single-script account, use an ID like `zone:ae36f999674d196762efcc5abb06b345`

```
$ terraform import cloudflare_worker_script.default zone:ae36f999674d196762efcc5abb06b345
```

where:

* `ae36f999674d196762efcc5abb06b345` - the zone ID

### multi-script

To import a script from a multi-script account, use an id like `name:script_name`

```
$ terraform import cloudflare_worker_script.default name:script_name
```

where:

* `script_name` - the script name
