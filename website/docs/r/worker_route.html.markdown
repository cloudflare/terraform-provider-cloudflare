---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_worker_route"
sidebar_current: "docs-cloudflare-resource-worker-route"
description: |-
  Provides a Cloudflare worker route resource.
---

# cloudflare_worker_route

Provides a Cloudflare worker route resource. A route will also require a `cloudflare_worker_script`. *NOTE:*  This resource uses the Cloudflare account APIs. This requires setting the `CLOUDFLARE_ACCOUNT_ID` environment variable or `account_id` provider argument.

## Example Usage

```hcl
# Runs the specified worker script for all URLs that match `example.com/*`
resource "cloudflare_worker_route" "my_route" {
  zone_id = "d41d8cd98f00b204e9800998ecf8427e"
  pattern = "example.com/*"
  script_name = cloudflare_worker_script.my_script.name
}

resource "cloudflare_worker_script" "my_script" {
  # see "cloudflare_worker_script" documentation ...
}
```

## Argument Reference

The following arguments are supported:

* `zone_id` - (Required) The zone ID to add the route to.
* `pattern` - (Required) The [route pattern](https://developers.cloudflare.com/workers/about/routes/)
* `script_name` Which worker script to run for requests that match the route pattern. If `script_name` is empty, workers will be skipped for matching requests.

## Import

Records can be imported using a composite ID formed of zone ID and route ID, e.g.

```
$ terraform import cloudflare_worker_route.default d41d8cd98f00b204e9800998ecf8427e/9a7806061c88ada191ed06f989cc3dac
```

where:

* `d41d8cd98f00b204e9800998ecf8427e` - zone ID
* `9a7806061c88ada191ed06f989cc3dac` - route ID as returned by [API](https://api.cloudflare.com/#worker-filters-list-filters)
