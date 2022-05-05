---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_workers_kv"
sidebar_current: "docs-cloudflare-resource-workers-kv"
description: |-
  Provides the ability to manage Cloudflare Workers KV features.
---

# cloudflare_workers_kv

Provides a Workers KV Pair.  *NOTE:*  This resource uses the Cloudflare account APIs.  This requires setting the `CLOUDFLARE_ACCOUNT_ID` environment variable or `account_id` provider argument.

## Example Usage

```hcl
resource "cloudflare_workers_kv_namespace" "example_ns" {
  title = "test-namespace"
}

resource "cloudflare_workers_kv" "example" {
  namespace_id = cloudflare_workers_kv_namespace.example_ns.id
  key = "test-key"
  value = "test value"
}
```

## Argument Reference

The following arguments are supported:

* `namespace_id` - (Required) The ID of the Workers KV namespace in which you want to create the KV pair
* `key` - (Required) The key name
* `value` - (Required) The string value to be stored in the key


## Import

Workers KV Namespace settings can be imported using it's ID.  **Note:** Because the same key can exist in multiple namespaces, the ID is a composite key of the format <namespace_id>/<key>.

```
$ terraform import cloudflare_workers_kv.example beaeb6716c9443eaa4deef11763ccca6/test-key
```

where:
- `beaeb6716c9443eaa4deef11763ccca6` is the ID of the namespace and `test-key` is the key
