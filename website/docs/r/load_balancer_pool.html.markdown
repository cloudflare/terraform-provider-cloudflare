---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_load_balancer_pool"
sidebar_current: "docs-cloudflare-resource-load-balancer-pool"
description: |-
  Provides a Cloudflare Load Balancer Pool resource.
---

# cloudflare_load_balancer_pool

Provides a Cloudflare Load Balancer pool resource. This provides a pool of origins that can be used by a Cloudflare Load Balancer. Note that the load balancing feature must be enabled in your Clouflare account before you can use this resource.


## Example Usage

```hcl
resource "cloudflare_load_balancer_pool" "foo" {
  name = "example-pool"
  origins {
    name = "example-1"
    address = "192.0.2.1"
    enabled = false
  }
  origins {
    name = "example-2"
    address = "192.0.2.2"
  }
  description = "example load balancer pool"
  enabled = false
  minimum_origins = 1
  notification_email = "someone@example.com"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) A short name (tag) for the pool. Only alphanumeric characters, hyphens and underscores are allowed.
* `origins` - (Required) The list of origins within this pool. Traffic directed at this pool is balanced across all currently healthy origins, provided the pool itself is healthy. Fields documented below
* `check_regions` - (Optional) A list of regions from which to run health checks. Empty means every Cloudflare datacenter (the default).
* `description` - (Optional) Free text description.
* `enabled` - (Optional) Whether to enable (the default) this pool. Disabled pools will not receive traffic and are excluded from health checks. Disabling a pool will cause any load balancers using it to failover to the next pool (if any).
* `minimum_origins` - (Optional) The minimum number of origins that must be healthy for this pool to serve traffic. If the number of healthy origins falls below this number, the pool will be marked unhealthy and we wil failover to the next available pool. Default: 1.
* `monitor` - (Optional) The ID of the Monitor to use for health checking origins within this pool.
* `notification_email` - (Optional) The email address to send health status notifications to. This can be an individual mailbox or a mailing list.

## Attributes Reference

The following attributes are exported:

* `id` - ID for this load balancer pool.
* `created_on` - The RFC3339 timestamp of when the load balancer was created.
* `modified_on` - The RFC3339 timestamp of when the load balancer was last modified.
