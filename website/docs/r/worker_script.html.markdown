---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_worker_script"
sidebar_current: "docs-cloudflare-resource-worker-script"
description: |-
  Provides a Cloudflare worker script resource.
---

# cloudflare_worker_script

Provides a Cloudflare worker script resource. In order for a script to be active, you'll also need to setup a `cloudflare_worker_route`. *NOTE:*  This resource uses the Cloudflare account APIs. This requires setting the `CLOUDFLARE_ACCOUNT_ID` environment variable or `account_id` provider argument.

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
    name = "DB"
    namespace_id = cloudflare_workers_kv_namespace.my_namespace.id
  }

  plain_text_binding {
    name = "ENVIRONMENT"
    text = "staging"
  }

  secret_text_binding {
    name = "SENTRY_DSN"
    text = var.sentry_dsn
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name for the script.
* `content` - (Required) The script content.

**kv_namespace_binding** supports:

* `name` - (Required) The global variable for the binding in your Worker code.
* `kv_namespace_id` - (Required) ID of the KV namespace you want to use.

**plain_text_binding** supports:

* `name` - (Required) The global variable for the binding in your Worker code.
* `text` - (Required) The plain text you want to store.

**secret_text_binding** supports:

* `name` - (Required) The global variable for the binding in your Worker code.
* `text` - (Required) The secret text you want to store.

## Import

To import a script, use a script name, e.g. `script_name`

```
$ terraform import cloudflare_worker_script.default script_name
```

where:

* `script_name` - the script name
