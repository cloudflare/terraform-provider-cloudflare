---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_load_balancer_pools"
sidebar_current: "docs-cloudflare-datasource-load-balancer-pools"
description: |-
  List configured Cloudflare Load Balancer Pools.
---

# cloudflare_load_balancer_pools

Use this data source to look up [Load Balancer Pools][1].

## Example Usage

The example below list all configured Load Balancer Pools. The list is then returned as output.

```hcl
data "cloudflare_load_balancer_pools" "test" {
}

output "lb_pools" {
  value = data.cloudflare_load_balancer_pools.test.pools
}
```

## Attributes Reference

* `pools` - A list of Load Balancer Pools details. Full list below:

**pools**

* `id` - ID for this load balancer pool.
* `name` - Short name (tag) for the pool.
* `origins` - The list of origins within this pool. It's a complex value. See description below.
* `check_regions` - List of regions (specified by region code) from which to run health checks. Empty means every Cloudflare data center (the default), but requires an Enterprise plan. Region codes can be found [here](https://support.cloudflare.com/hc/en-us/articles/115000540888-Load-Balancing-Geographic-Regions).
* `description` - Text description.
* `load_shedding` - Setting for controlling load shedding for this pool.
* `enabled` - Whether this pool is enabled. Disabled pools will not receive traffic and are excluded from health checks.
* `minimum_origins` Minimum number of origins that must be healthy for this pool to serve traffic.
* `monitor` - ID of the Monitor to use for health checking origins within this pool.
* `notification_email` - Email address to send health status notifications to. Multiple emails are set as a comma delimited list.
* `latitude` - Latitude this pool is physically located at; used for proximity steering.
* `longitude` - Longitude this pool is physically located at; used for proximity steering.
* `created_on` - The RFC3339 timestamp of when the load balancer was created.
* `modified_on` - The RFC3339 timestamp of when the load balancer was last modified.

**origins**:

* `name` - A human-identifiable name for the origin.
* `address` - The IP address (IPv4 or IPv6) of the origin, or the publicly addressable hostname.
* `weight` - The weight (0.01 - 1.00) of this origin, relative to other origins in the pool. Equal values mean equal weighting. A weight of 0 means traffic will not be sent to this origin, but health is still checked.
* `enabled` - Whether this origin is enabled. Disabled origins will not receive traffic and are excluded from health checks.
* `header` - HTTP request headers. For security reasons, this header also needs to be a subdomain of the overall zone. Fields documented below.

**load_shedding**:
* `default_percent` - Percent of traffic to shed 0 - 100.
* `default_policy` - Method of shedding traffic "", "hash" or "random".
* `session_percent` - Percent of session traffic to shed 0 - 100.
* `session_policy` - Method of shedding session traffic "" or "hash".

**header**

* `header` - Header name.
* `values` - List of string values for the header.

[1]: https://api.cloudflare.com/#load-balancer-pools-properties
