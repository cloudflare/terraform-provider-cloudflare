---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_load_balancer_pool"
sidebar_current: "docs-cloudflare-resource-load-balancer-pool"
description: |-
  Provides a Cloudflare Load Balancer Pool resource.
---

# cloudflare_load_balancer_pool

Provides a Cloudflare Load Balancer pool resource. This provides a pool of origins that can be used by a Cloudflare Load Balancer. Note that the load balancing feature must be enabled in your Cloudflare account before you can use this resource.


## Example Usage

```hcl
resource "cloudflare_load_balancer_pool" "foo" {
  name = "example-pool"
  origins {
    name = "example-1"
    address = "192.0.2.1"
    enabled = false
    header {
      header = "Host"
      values = ["example-1"]
    }
  }
  origins {
    name = "example-2"
    address = "192.0.2.2"
    header {
      header = "Host"
      values = ["example-2"]
    }
  }
  description = "example load balancer pool"
  enabled = false
  minimum_origins = 1
  notification_email = "someone@example.com"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) A short name (tag) for the pool. Only alphanumeric characters, hyphens, and underscores are allowed.
* `origins` - (Required) The list of origins within this pool. Traffic directed at this pool is balanced across all currently healthy origins, provided the pool itself is healthy. It's a complex value. See description below.
* `check_regions` - (Optional) A list of regions (specified by region code) from which to run health checks. Empty means every Cloudflare data center (the default), but requires an Enterprise plan. Region codes can be found [here](https://support.cloudflare.com/hc/en-us/articles/115000540888-Load-Balancing-Geographic-Regions).
* `description` - (Optional) Free text description.
* `enabled` - (Optional) Whether to enable (the default) this pool. Disabled pools will not receive traffic and are excluded from health checks. Disabling a pool will cause any load balancers using it to failover to the next pool (if any).
* `minimum_origins` - (Optional) The minimum number of origins that must be healthy for this pool to serve traffic. If the number of healthy origins falls below this number, the pool will be marked unhealthy and we will failover to the next available pool. Default: 1.
* `monitor` - (Optional) The ID of the Monitor to use for health checking origins within this pool.
* `notification_email` - (Optional) The email address to send health status notifications to. This can be an individual mailbox or a mailing list. Multiple emails can be supplied as a comma delimited list.

The **origins** block supports:

* `name` - (Required) A human-identifiable name for the origin.
* `address` - (Required) The IP address (IPv4 or IPv6) of the origin, or the publicly addressable hostname. Hostnames entered here should resolve directly to the origin, and not be a hostname proxied by Cloudflare.
* `weight` - (Optional) The weight (0.01 - 1.00) of this origin, relative to other origins in the pool. Equal values mean equal weighting. A weight of 0 means traffic will not be sent to this origin, but health is still checked. Default: 1.
* `enabled` - (Optional) Whether to enable (the default) this origin within the Pool. Disabled origins will not receive traffic and are excluded from health checks. The origin will only be disabled for the current pool.
* `header` - (Optional)  The HTTP request headers. For security reasons, this header also needs to be a subdomain of the overall zone. Fields documented below.

**header** requires the following:

* `header` - (Required) The header name.
* `values` - (Required) A list of string values for the header.

## Attributes Reference

The following attributes are exported:

* `id` - ID for this load balancer pool.
* `created_on` - The RFC3339 timestamp of when the load balancer was created.
* `modified_on` - The RFC3339 timestamp of when the load balancer was last modified.
