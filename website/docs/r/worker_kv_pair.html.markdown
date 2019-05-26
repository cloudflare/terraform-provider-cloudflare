---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_worker_kv_pair"
sidebar_current: "docs-cloudflare-resource-worker-kv-pair"
description: |-
  Provides a Cloudflare worker kv pair resource.
---

# cloudflare_worker_kv_pair

Provides a Cloudflare worker kv pair resource. A pair will require both a `cloudflare_worker_kv_namespace` and a `cloudflare_worker_script`.

## Example Usage

```hcl
# Create a new key-value pair from a local file
resource "cloudflare_worker_kv_pair" "my_pair" {
  namespace = "${cloudflare_worker_kv_namespace.my_namespace.id}"
  key = "my-key"
  value = "${my-file.txt}"

  # it's recommended to set `depends_on` to make sure that the
  # kv namespace and script have both been created
  depends_on = ["cloudflare_worker_kv_namespace.my_namespace"]
}

resource "cloudflare_worker_kv_namespace" "my_namespace" {
  # see "cloudflare_worker_kv_namespace" documentation ...
  depends_on = ["cloudflare_worker_kv_namespace.my_script"]
}

resource "cloudflare_worker_script" "my_script" {
  # see "cloudflare_worker_script" documentation ...
}
```

## Argument Reference

The following arguments are supported:

* `namespace` - (Required) The ID of the worker kv namespace.
* `key` - (Required) The key name of the pair.
* `value` - (Required) The value of the pair.

## Import

Records can be imported using the ID of the pair, e.g.

```
$ terraform import cloudflare_worker_kv_pair.default 0f2ac74b498b48028cb68387c421e279/my-key
```

where:

* `0f2ac74b498b48028cb68387c421e279` - namespace ID as returned by [API](https://api.cloudflare.com/#workers-kv-namespace-create-a-namespace)
* `my-key` - key name of the pair


