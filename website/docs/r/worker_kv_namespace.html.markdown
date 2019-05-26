---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_worker_kv_namespace"
sidebar_current: "docs-cloudflare-resource-worker-kv-namespace"
description: |-
  Provides a Cloudflare worker kv namespace resource.
---

# cloudflare_worker_kv_namespace

Provides a Cloudflare worker kv namespace resource. A namespace will also require a `cloudflare_worker_script`.

## Example Usage

```hcl
# Creates a new workers kv namespace titled 'My Namespace'
resource "cloudflare_worker_kv_namespace" "my_namespace" {
  title = "My Namespace"

  # it's recommended to set `depends_on` to point to the cloudflare_worker_script
  # resource in order to make sure that the script is uploaded before the route
  # is created
  depends_on = ["cloudflare_worker_script.my_script"]
}

resource "cloudflare_worker_script" "my_script" {
  # see "cloudflare_worker_script" documentation ...
}
```

## Argument Reference

The following arguments are supported:

* `title` - (Required) The human-readable title of the namespace.

## Import

Records can be imported using the ID of the namespace, e.g.

```
$ terraform import cloudflare_worker_kv_namespace.default 0f2ac74b498b48028cb68387c421e279
```

where:

* `0f2ac74b498b48028cb68387c421e279` - route ID as returned by [API](https://api.cloudflare.com/#workers-kv-namespace-create-a-namespace)


