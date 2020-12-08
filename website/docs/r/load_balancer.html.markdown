---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_load_balancer"
sidebar_current: "docs-cloudflare-resource-load-balancer"
description: |-
  Provides a Cloudflare Load Balancer resource.
---

# cloudflare_load_balancer

Provides a Cloudflare Load Balancer resource. This sits in front of a number of defined pools of origins and provides various options for geographically-aware load balancing. Note that the load balancing feature must be enabled in your Cloudflare account before you can use this resource.

## Example Usage

```hcl
# Define a load balancer which always points to a pool we define below
# In normal usage, would have different pools set for different pops (cloudflare points-of-presence) and/or for different regions
# Within each pop or region we can define multiple pools in failover order
resource "cloudflare_load_balancer" "bar" {
  zone_id = "d41d8cd98f00b204e9800998ecf8427e"
  name = "example-load-balancer.example.com"
  fallback_pool_id = cloudflare_load_balancer_pool.foo.id
  default_pool_ids = [cloudflare_load_balancer_pool.foo.id]
  description = "example load balancer using geo-balancing"
  proxied = true
  steering_policy = "geo"
  pop_pools {
    pop = "LAX"
    pool_ids = [cloudflare_load_balancer_pool.foo.id]
  }
  region_pools {
    region = "WNAM"
    pool_ids = [cloudflare_load_balancer_pool.foo.id]
  }
}

resource "cloudflare_load_balancer_pool" "foo" {
  name = "example-lb-pool"
  origins {
    name = "example-1"
    address = "192.0.2.1"
    enabled = false
  }
}
```

## Argument Reference

The following arguments are supported:

* `zone_id` - (Required) The zone ID to add the load balancer to.
* `name` - (Required) The DNS name (FQDN, including the zone) to associate with the load balancer.
* `fallback_pool_id` - (Required) The pool ID to use when all other pools are detected as unhealthy.
* `default_pool_ids` - (Required) A list of pool IDs ordered by their failover priority. Used whenever region/pop pools are not defined.
* `description` - (Optional) Free text description.
* `ttl` - (Optional) Time to live (TTL) of this load balancer's DNS `name`. Conflicts with `proxied` - this cannot be set for proxied load balancers. Default is `30`.
* `steering_policy` - (Optional) Determine which method the load balancer uses to determine the fastest route to your origin. Valid values are: `"off"`, `"geo"`, `"dynamic_latency"`, `"random"` or `""`. Default is `""`.
* `proxied` - (Optional) Whether the hostname gets Cloudflare's origin protection. Defaults to `false`.
* `enabled` - (Optional) Enable or disable the load balancer. Defaults to `true` (enabled).
* `region_pools` - (Optional) A set containing mappings of region/country codes to a list of pool IDs (ordered by their failover priority) for the given region. Fields documented below.
* `pop_pools` - (Optional) A set containing mappings of Cloudflare Point-of-Presence (PoP) identifiers to a list of pool IDs (ordered by their failover priority) for the PoP (datacenter). This feature is only available to enterprise customers. Fields documented below.
* `session_affinity` - (Optional) Associates all requests coming from an end-user with a single origin. Cloudflare will set a cookie on the initial response to the client, such that consequent requests with the cookie in the request will go to the same origin, so long as it is available.  Valid values are: `""`, `"none"`, `"cookie"`, and `"ip_cookie"`.  Default is `""`.
* `session_affinity_ttl` - (Optional) Time, in seconds, until this load balancers session affinity cookie expires after being created. This parameter is ignored unless a supported session affinity policy is set. The current default of 23 hours will be used unless `session_affinity_ttl` is explicitly set. Once the expiry time has been reached, subsequent requests may get sent to a different origin server. Valid values are between 1800 and 604800.

**region_pools** requires the following:

* `region` - (Required) A region code which must be in the list defined [here](https://support.cloudflare.com/hc/en-us/articles/115000540888-Load-Balancing-Geographic-Regions). Multiple entries should not be specified with the same region.
* `pool_ids` - (Required) A list of pool IDs in failover priority to use in the given region.

**pop_pools** requires the following:

* `pop` - (Required) A 3-letter code for the Point-of-Presence. Allowed values can be found in the list of datacenters on the [status page](https://www.cloudflarestatus.com/). Multiple entries should not be specified with the same PoP.
* `pool_ids` - (Required) A list of pool IDs in failover priority to use for traffic reaching the given PoP.

## Attributes Reference

The following attributes are exported:

* `id` - Unique identifier in the API for the load balancer.
* `created_on` - The RFC3339 timestamp of when the load balancer was created.
* `modified_on` - The RFC3339 timestamp of when the load balancer was last modified.
