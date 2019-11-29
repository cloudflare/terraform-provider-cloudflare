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

```hcl
resource "cloudflare_workers_kv_namespace" "my_namespace" {
  title = "example"
}

# Sets the script with the name "script_1"
resource "cloudflare_worker_script" "my_script" {
  name = "script_1"
  content = file("script.js")
  
  kv_namespace_binding {
    name = "my_binding"
    namespace_id = cloudflare_workers_kv_namespace.my_namespace.id
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name for the script.
* `content` - (Required) The script content.

**kv_namespace_binding** (optional) block supports:

* `name` - (Required) The name for the binding.
* `namespace_id` - (Required) ID of KV namespace.

## Import

To import a script, use a script name, e.g. `script_name`

```
$ terraform import cloudflare_worker_script.default script_name
```

where:

* `script_name` - the script name
