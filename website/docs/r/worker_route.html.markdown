---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_worker_route"
sidebar_current: "docs-cloudflare-resource-worker-route"
description: |-
  Provides a Cloudflare worker route resource.
---

# cloudflare_worker_route

Provides a Cloudflare worker route resource. A route will also require a `cloudflare_worker_script`.

## Example Usage

__NOTE:__ This is for non-enterprise accounts where there is one script per zone. The `enabled` flag determines whether to run the worker script for a request that matches the specified `pattern`. For enterprise accounts, see the "multi-script" example below.

```hcl
# Enables the zone's worker script for all URLs that match `example.com/*`
resource "cloudflare_worker_route" "my_route" {
  zone = "example.com"
  pattern = "example.com/*"
  enabled = true

  # it's recommended to set `depends_on` to point to the cloudflare_worker_script
  # resource in order to make sure that the script is uploaded before the route
  # is created
  depends_on = ["cloudflare_worker_script.my_script"]
}

resource "cloudflare_worker_script" "my_script" {
  # see "cloudflare_worker_script" documentation ...
}
```

## Multi-script example usage

__NOTE:__ This is only for enterprise accounts. With multi-script, each route points to a particular script instead of setting an `enabled` flag

```hcl
# Runs the specified worker script for all URLs that match `example.com/*`
resource "cloudflare_worker_route" "my_route" {
  zone = "example.com"
  pattern = "example.com/*"
  script_name = "${cloudflare_worker_script.my_script.name}"
}

resource "cloudflare_worker_script" "my_script" {
  # see "cloudflare_worker_script" documentation ...
}
```

## Argument Reference

The following arguments are supported:

* `zone` - (Required) The zone to add the route to.
* `pattern` - (Required) The [route pattern](https://developers.cloudflare.com/workers/api/route-matching/)
* `enabled` (For single-script accounts only) Whether to run the worker script for requests that match the route pattern. Default is `false`
* `script_name` (For multi-script accounts only) Which worker script to run for requests that match the route pattern. If `script_name` is empty, workers will be skipped for matching requests.

## Attributes Reference

The following attributes are exported:

* `zone_id` - The zone id of the route

## Import

Records can be imported using a composite ID formed of zone name and route ID, e.g.

```
$ terraform import cloudflare_worker_route.default example.com/9a7806061c88ada191ed06f989cc3dac
```

where:

* `9a7806061c88ada191ed06f989cc3dac` - route ID as returned by [API](https://api.cloudflare.com/#worker-filters-list-filters)


