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
  zone = "example.com"
  content = "${file("script.js")}"
}
```

## Multi-script example usage

__NOTE:__ This is only for enterprise accounts. With multi-script, each script is given a `name` instead of a `zone`

```hcl
# Sets the script with the name "script_1"
resource "cloudflare_worker_script" "my_script" {
  name = "script_1"
  content = "${file("script.js")}"
}
```

## Argument Reference

The following arguments are supported:

* `zone` - (Required for single-script accounts) The zone for the script.
* `name` - (Required for multi-script accounts) The name for the script. 
* `content` - (Required) The script content.

## Attributes Reference

The following attributes are exported:

* `zone_id` - The zone id of the script (only for non-multi-script resources)

## Import

### single-script

To import a script from a single-script account, use an id like `zone:example.com`

```
$ terraform import cloudflare_worker_script.default zone:example.com
```

where:

* `example.com` - the zone name

### multi-script

To import a script from a multi-script account, use an id like `name:script_name`

```
$ terraform import cloudflare_worker_script.default name:script_name
```

where:

* `script_name` - the script name


