---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_load_balancer"
sidebar_current: "docs-cloudflare-resource-load-balancer"
description: |-
  Provides a Cloudflare Load Balancer resource.
---

# cloudflare_load_balancer

Provides a Cloudflare Load Balancer resource. This sits in front of a number of defined pools of origins and provides various options for geographically-aware load balancing.

## Example Usage

```hcl
resource "cloudflare_load_balancer_pool" "foo" {
  name = "example-lb-pool"
  origins {
    name = "example-1"
    address = "192.0.2.1"
    enabled = false
  }
}

# Define a load balancer which always points to the pool we defined
# In normal usage, would have different pools set for different pops/regions
# And some failover defined within that
resource "cloudflare_load_balancer" "bar" {
  zone = "example.com"
  name = "example-load-balancer"
  fallback_pool_id = "${cloudflare_load_balancer_pool.foo.id}"
  default_pool_ids = ["${cloudflare_load_balancer_pool.foo.id}"]
  description = "example load balancer using geo-balancing"
  proxied = true
  pop_pools {
    pop = "LAX"
    pool_ids = ["${cloudflare_load_balancer_pool.foo.id}"]
  }
  region_pools {
    region = "WNAM"
    pool_ids = ["${cloudflare_load_balancer_pool.foo.id}"]
  }
}
```

## Argument Reference

The following arguments are supported:

* `zone` - (Required) The zone to add the load balancer to.
* `name` - (Required) The DNS name to associate with the load balancer.
* `fallback_pool_id` - (Required) The pool ID to use when all other pools are detected as unhealthy.
* `default_pool_ids` - (Required) A list of pool IDs ordered by their failover priority. Used whenever region/pop pools are not defined.
* `description` - (Optional) Free text description.
* `ttl` - (Optional) Time to live (TTL) of this load balancer's DNS `name`. Conflicts with `proxied` - this cannot be set for proxied load balancers. Default is `30`.
* `proxied` - (Optional) Whether the hostname gets Cloudflare's origin protection. Defaults to `false`.
* `region_pools` - (Optional) A set containing mappings of region/country codes to a list of pool IDs (ordered by their failover priority) for the given region. Fields documented below.
* `pop_pools` - (Optional) A set containing mappings of Cloudflare Point-of-Presence identifiers to a list of pool IDs (ordered by their failover priority) for the PoP (datacenter). This feature is only available to enterprise customers. Fields documented below.

**region_pools** requires the following:

* `region` - (Required) A region code which must be in the list defined [here](https://support.cloudflare.com/hc/en-us/articles/115000540888-Load-Balancing-Geographic-Regions).
* `pool_ids` - (Required) A list of pool IDs in failover priority to use in the given region.

**pop_pools** requires the following:

* `pop` - (Required) A 3-letter code for the PoP. Allowed values can be found in the list of datacenters on the [status page](https://www.cloudflarestatus.com/).
* `pool_ids` - (Required) A list of pool IDs in failover priority to use for traffic reaching the given PoP.

## Attributes Reference

The following attributes are exported:

* `id` - Unique identifier in the API for the load balancer.
* `zone_id` - ID associated with the specified `zone`.
* `created_on` - The RFC3339 timestamp of when the load balancer was created.
* `modified_on` - The RFC3339 timestamp of when the load balancer was last modified.