---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_workers_kv_namespace"
sidebar_current: "docs-cloudflare-resource-workers-kv-namespace"
description: |-
  Provides the ability to manage Cloudflare Workers KV Namespace features.
---

# cloudflare_workers_kv_namespace

Provides a Workers KV Namespace

## Example Usage

```hcl
resource "cloudflare_workers_kv_namespace" "example" {
  title        = "test-namespace"
}
```

## Argument Reference

The following arguments are supported:

* `title` - (Required) The name of the namespace you wish to create


## Import

KV Namespace settings can be imported using the it's title e.g.

```
$ terraform import cloudflare_kv_namespace.example title
```

where `title` is the title of the namespace
