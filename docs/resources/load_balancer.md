---
page_title: "cloudflare_load_balancer Resource - Cloudflare"
subcategory: ""
description: |-
  Accepted Permissions
  Load Balancers ReadLoad Balancers Write
---

# cloudflare_load_balancer (Resource)

Accepted Permissions

- `Load Balancers Read`
- `Load Balancers Write`

## Planning updates

On an update, unchanged API defaults can retain their state values instead of
showing `(known after apply)`. This includes unset overrides in unchanged, named
rules. Removing a configured, non-default value still lets the API reset that
value; the plan does not retain the removed setting.

`modified_on` remains unknown because it changes on every write. Other computed
values can remain unknown on creation, when a related setting changes, or when
the provider cannot determine a safe default. For example, changing session
affinity can change its default TTL, and an unconfigured DNS TTL remains unknown
for an unproxied load balancer.

## Example Usage

```terraform
resource "cloudflare_load_balancer" "example_load_balancer" {
  default_pools = ["17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"]
  fallback_pool = "fallback_pool"
  name = "www.example.com"
  zone_id = "zone_id"
  adaptive_routing = {
    failover_across_pools = true
  }
  country_pools = {
    GB = ["abd90f38ced07c2e2f4df50b1f61d4194"]
    US = ["de90f38ced07c2e2f4df50b1f61d4194", "00920f38ce07c2e2f4df50b1f61d4194"]
  }
  description = "Load Balancer for www.example.com"
  enabled = true
  location_strategy = {
    mode = "resolver_ip"
    prefer_ecs = "always"
  }
  networks = ["cloudflare", "jdcloud"]
  pool_sets = [{
    disabled = true
    fixed_response = {
      content_type = "application/json"
      location = "www.example.com"
      message_body = "Testing Hello"
      status_code = 0
    }
    match = {
      default = true
      topology = {
        countries = ["string"]
        pops = ["string"]
        regions = ["WNAM"]
      }
    }
    name = "wnam-primary"
    overrides = {
      fallback_pool = "9290f38c5d07c2e2f4df57b1f61d4196"
      pool_default_weight = 0.2
      pool_weights = {
        "9290f38c5d07c2e2f4df57b1f61d4196" = 0.5
        de90f38ced07c2e2f4df50b1f61d4194 = 0.3
      }
      pools = ["17b5962d775c646f3f9725cbc7a53df4"]
      steering_policy = "random"
    }
  }]
  pop_pools = {
    LAX = ["de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"]
    LHR = ["abd90f38ced07c2e2f4df50b1f61d4194", "f9138c5d07c2e2f4df57b1f61d4196"]
    SJC = ["00920f38ce07c2e2f4df50b1f61d4194"]
  }
  proxied = true
  random_steering = {
    default_weight = 0.2
    pool_weights = {
      "9290f38c5d07c2e2f4df57b1f61d4196" = 0.5
      de90f38ced07c2e2f4df50b1f61d4194 = 0.3
    }
  }
  region_pools = {
    ENAM = ["00920f38ce07c2e2f4df50b1f61d4194"]
    WNAM = ["de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"]
  }
  rules = [{
    condition = "http.request.uri.path contains \"/testing\""
    disabled = true
    fixed_response = {
      content_type = "application/json"
      location = "www.example.com"
      message_body = "Testing Hello"
      status_code = 0
    }
    name = "route the path /testing to testing datacenter."
    overrides = {
      adaptive_routing = {
        failover_across_pools = true
      }
      country_pools = {
        GB = ["abd90f38ced07c2e2f4df50b1f61d4194"]
        US = ["de90f38ced07c2e2f4df50b1f61d4194", "00920f38ce07c2e2f4df50b1f61d4194"]
      }
      default_pools = ["17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"]
      fallback_pool = "fallback_pool"
      location_strategy = {
        mode = "resolver_ip"
        prefer_ecs = "always"
      }
      pool_default_weight = 0.2
      pool_weights = {
        "9290f38c5d07c2e2f4df57b1f61d4196" = 0.5
        de90f38ced07c2e2f4df50b1f61d4194 = 0.3
      }
      pools = ["17b5962d775c646f3f9725cbc7a53df4"]
      pop_pools = {
        LAX = ["de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"]
        LHR = ["abd90f38ced07c2e2f4df50b1f61d4194", "f9138c5d07c2e2f4df57b1f61d4196"]
        SJC = ["00920f38ce07c2e2f4df50b1f61d4194"]
      }
      random_steering = {
        default_weight = 0.2
        pool_weights = {
          "9290f38c5d07c2e2f4df57b1f61d4196" = 0.5
          de90f38ced07c2e2f4df50b1f61d4194 = 0.3
        }
      }
      region_pools = {
        ENAM = ["00920f38ce07c2e2f4df50b1f61d4194"]
        WNAM = ["de90f38ced07c2e2f4df50b1f61d4194", "9290f38c5d07c2e2f4df57b1f61d4196"]
      }
      session_affinity = "cookie"
      session_affinity_attributes = {
        drain_duration = 100
        headers = ["x"]
        require_all_headers = true
        samesite = "Auto"
        secure = "Auto"
        zero_downtime_failover = "sticky"
      }
      session_affinity_ttl = 1800
      steering_policy = "dynamic_latency"
      ttl = 30
    }
    priority = 0
    terminates = true
  }]
  session_affinity = "cookie"
  session_affinity_attributes = {
    drain_duration = 100
    headers = ["x"]
    require_all_headers = true
    samesite = "Auto"
    secure = "Auto"
    zero_downtime_failover = "sticky"
  }
  session_affinity_ttl = 1800
  steering_policy = "dynamic_latency"
  ttl = 30
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `default_pools` (List of String) A list of pool IDs ordered by their failover priority. Pools defined here are used by default, or when region_pools are not configured for a given region.
- `fallback_pool` (String) The pool ID to use when all other pools are detected as unhealthy.
- `name` (String) The DNS hostname to associate with your Load Balancer. If this hostname already exists as a DNS record in Cloudflare's DNS, the Load Balancer will take precedence and the DNS record will not be used.
- `zone_id` (String)

### Optional

- `adaptive_routing` (Attributes) Controls features that modify the routing of requests to pools and origins in response to dynamic conditions, such as during the interval between active health monitoring requests. For example, zero-downtime failover occurs immediately when an origin becomes unavailable due to HTTP 521, 522, or 523 response codes. If there is another healthy origin in the same pool, the request is retried once against this alternate origin. (see [below for nested schema](#nestedatt--adaptive_routing))
- `country_pools` (Map of List of String) A mapping of country codes to a list of pool IDs (ordered by their failover priority) for the given country. Any country not explicitly defined will fall back to using the corresponding region_pool mapping if it exists else to default_pools.
- `description` (String) Object description.
- `enabled` (Boolean) Whether to enable (the default) this load balancer.
- `location_strategy` (Attributes) Controls location-based steering for non-proxied requests. See `steering_policy` to learn how steering is affected. (see [below for nested schema](#nestedatt--location_strategy))
- `networks` (List of String) List of networks where Load Balancer or Pool is enabled.
- `pool_sets` (Attributes List) An optional list of pool sets, evaluated in array order with first match wins. Pool sets are independent from the standard steering fields (`region_pools` / `country_pools` / `pop_pools` / `default_pools` / `steering_policy` / `random_steering` / `fallback_pool` / `rules`). On a PATCH, an empty array (`pool_sets: []`) clears all pool sets, while omitting the field leaves existing pool sets unchanged. (see [below for nested schema](#nestedatt--pool_sets))
- `pop_pools` (Map of List of String) Enterprise only: A mapping of Cloudflare PoP identifiers to a list of pool IDs (ordered by their failover priority) for the PoP (datacenter). Any PoPs not explicitly defined will fall back to using the corresponding country_pool, then region_pool mapping if it exists else to default_pools.
- `proxied` (Boolean) Whether the hostname should be gray clouded (false) or orange clouded (true).
- `random_steering` (Attributes) Configures pool weights.
- `steering_policy="random"`: A random pool is selected with probability proportional to pool weights.
- `steering_policy="least_outstanding_requests"`: Use pool weights to scale each pool's outstanding requests.
- `steering_policy="least_connections"`: Use pool weights to scale each pool's open connections. (see [below for nested schema](#nestedatt--random_steering))
- `region_pools` (Map of List of String) A mapping of region codes to a list of pool IDs (ordered by their failover priority) for the given region. Any regions not explicitly defined will fall back to using default_pools.
- `rules` (Attributes List) BETA Field Not General Access: A list of rules for this load balancer to execute. (see [below for nested schema](#nestedatt--rules))
- `session_affinity` (String) Specifies the type of session affinity the load balancer should use unless specified as `"none"`. The supported types are: - `"cookie"`: On the first request to a proxied load balancer, a cookie is generated, encoding information of which origin the request will be forwarded to. Subsequent requests, by the same client to the same load balancer, will be sent to the origin server the cookie encodes, for the duration of the cookie and as long as the origin server remains healthy. If the cookie has expired or the origin server is unhealthy, then a new origin server is calculated and used. - `"ip_cookie"`: Behaves the same as `"cookie"` except the initial origin selection is stable and based on the client's ip address. - `"header"`: On the first request to a proxied load balancer, a session key based on the configured HTTP headers (see `session_affinity_attributes.headers`) is generated, encoding the request headers used for storing in the load balancer session state which origin the request will be forwarded to. Subsequent requests to the load balancer with the same headers will be sent to the same origin server, for the duration of the session and as long as the origin server remains healthy. If the session has been idle for the duration of `session_affinity_ttl` seconds or the origin server is unhealthy, then a new origin server is calculated and used. See `headers` in `session_affinity_attributes` for additional required configuration.
Available values: "none", "cookie", "ip_cookie", "header".
- `session_affinity_attributes` (Attributes) Configures attributes for session affinity. (see [below for nested schema](#nestedatt--session_affinity_attributes))
- `session_affinity_ttl` (Number) Time, in seconds, until a client's session expires after being created. Once the expiry time has been reached, subsequent requests may get sent to a different origin server. The accepted ranges per `session_affinity` policy are: - `"cookie"` / `"ip_cookie"`: The current default of 23 hours will be used unless explicitly set. The accepted range of values is between [1800, 604800]. - `"header"`: The current default of 1800 seconds will be used unless explicitly set. The accepted range of values is between [30, 3600]. Note: With session affinity by header, sessions only expire after they haven't been used for the number of seconds specified.
- `steering_policy` (String) Steering Policy for this load balancer.
- `"off"`: Use `default_pools`.
- `"geo"`: Use `region_pools`/`country_pools`/`pop_pools`. For non-proxied requests, the country for `country_pools` is determined by `location_strategy`.
- `"random"`: Select a pool randomly.
- `"dynamic_latency"`: Use round trip time to select the closest pool in default_pools (requires pool health checks).
- `"proximity"`: Use the pools' latitude and longitude to select the closest pool using the Cloudflare PoP location for proxied requests or the location determined by `location_strategy` for non-proxied requests.
- `"least_outstanding_requests"`: Select a pool by taking into consideration `random_steering` weights, as well as each pool's number of outstanding requests. Pools with more pending requests are weighted proportionately less relative to others.
- `"least_connections"`: Select a pool by taking into consideration `random_steering` weights, as well as each pool's number of open connections. Pools with more open connections are weighted proportionately less relative to others. Supported for HTTP/1 and HTTP/2 connections.
- `""`: Will map to `"geo"` if you use `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
Available values: "off", "geo", "random", "dynamic_latency", "proximity", "least_outstanding_requests", "least_connections", "".
- `ttl` (Number) Time to live (TTL) of the DNS entry for the IP address returned by this load balancer. This only applies to gray-clouded (unproxied) load balancers.

### Read-Only

- `created_on` (String)
- `id` (String) The ID of this resource.
- `modified_on` (String)
- `zone_name` (String)

<a id="nestedatt--adaptive_routing"></a>
### Nested Schema for `adaptive_routing`

Optional:

- `failover_across_pools` (Boolean) Extends zero-downtime failover of requests to healthy origins from alternate pools, when no healthy alternate exists in the same pool, according to the failover order defined by traffic and origin steering. When set false (the default) zero-downtime failover will only occur between origins within the same pool. See `session_affinity_attributes` for control over when sessions are broken or reassigned.


<a id="nestedatt--location_strategy"></a>
### Nested Schema for `location_strategy`

Optional:

- `mode` (String) Determines the authoritative location when ECS is not preferred, does not exist in the request, or its GeoIP lookup is unsuccessful.
- `"pop"`: Use the Cloudflare PoP location.
- `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is unsuccessful, use the Cloudflare PoP location.
Available values: "pop", "resolver_ip".
- `prefer_ecs` (String) Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the authoritative location.
- `"always"`: Always prefer ECS.
- `"never"`: Never prefer ECS.
- `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
- `"geo"`: Prefer ECS only when `steering_policy="geo"`.
Available values: "always", "never", "proximity", "geo".


<a id="nestedatt--pool_sets"></a>
### Nested Schema for `pool_sets`

Optional:

- `disabled` (Boolean) Disable this specific pool set. It will no longer be evaluated.
- `fixed_response` (Attributes) A collection of fields used to directly respond to the client instead of routing to a pool. When supplied on a rule, that rule stops further rule evaluation. (see [below for nested schema](#nestedatt--pool_sets--fixed_response))
- `match` (Attributes) Determines which requests a pool set applies to. Set `topology` to match by location or `default: true` to match all requests; the two are mutually exclusive. A pool set with no `match` matches all requests. (see [below for nested schema](#nestedatt--pool_sets--match))
- `name` (String) A human-readable name for this pool set.
- `overrides` (Attributes) The behavior a pool set applies when its `match` succeeds. A strict subset of a rule's `overrides`: a pool set replaces the topology wholesale with a flat pool list (`pools`), so only the declarative pool-routing fields plus `fallback_pool` and `steering_policy` are settable. All fields are optional. (see [below for nested schema](#nestedatt--pool_sets--overrides))

<a id="nestedatt--pool_sets--fixed_response"></a>
### Nested Schema for `pool_sets.fixed_response`

Optional:

- `content_type` (String) The http 'Content-Type' header to include in the response.
- `location` (String) The http 'Location' header to include in the response.
- `message_body` (String) Text to include as the http body.
- `status_code` (Number) The http status code to respond with.


<a id="nestedatt--pool_sets--match"></a>
### Nested Schema for `pool_sets.match`

Optional:

- `default` (Boolean) When true, matches every request. Cannot be combined with `topology`.
- `topology` (Attributes) Matches requests by location. Set any combination of `pops`, `countries`, and `regions` (at least one is required); a request matches when its value appears in any populated list (e.g. `regions: ["WNAM"]` with `countries: ["US"]` matches a request in either WNAM or the US). (see [below for nested schema](#nestedatt--pool_sets--match--topology))

<a id="nestedatt--pool_sets--match--topology"></a>
### Nested Schema for `pool_sets.match.topology`

Optional:

- `countries` (List of String) A list of ISO 3166-1 alpha-2 country codes. Matches when the request's country is in this list.
- `pops` (List of String) A list of Cloudflare PoP codes. Matches when the request's PoP is in this list.
- `regions` (List of String) A list of Cloudflare region codes (e.g. `WNAM`, `ENAM`, `WEU`). Matches when the request's region is in this list.



<a id="nestedatt--pool_sets--overrides"></a>
### Nested Schema for `pool_sets.overrides`

Optional:

- `fallback_pool` (String) The pool ID to use when all other pools are detected as unhealthy.
- `pool_default_weight` (Number) The default weight for pools not listed in `pool_weights`. The declarative alternative to `random_steering.default_weight`; mutually exclusive with `random_steering`.
- `pool_weights` (Map of Number) A mapping of pool IDs to custom weights, relative to the other pools. The declarative alternative to `random_steering.pool_weights`; mutually exclusive with `random_steering`.
- `pools` (List of String) A flat, ordered list of pool IDs to route the matched audience to. Replaces the resolved topology with exactly these pools. Mutually exclusive with `fixed_response`.
- `steering_policy` (String) Steering Policy for this load balancer.
- `"off"`: Use `default_pools`.
- `"geo"`: Use `region_pools`/`country_pools`/`pop_pools`. For non-proxied requests, the country for `country_pools` is determined by `location_strategy`.
- `"random"`: Select a pool randomly.
- `"dynamic_latency"`: Use round trip time to select the closest pool in default_pools (requires pool health checks).
- `"proximity"`: Use the pools' latitude and longitude to select the closest pool using the Cloudflare PoP location for proxied requests or the location determined by `location_strategy` for non-proxied requests.
- `"least_outstanding_requests"`: Select a pool by taking into consideration `random_steering` weights, as well as each pool's number of outstanding requests. Pools with more pending requests are weighted proportionately less relative to others.
- `"least_connections"`: Select a pool by taking into consideration `random_steering` weights, as well as each pool's number of open connections. Pools with more open connections are weighted proportionately less relative to others. Supported for HTTP/1 and HTTP/2 connections.
- `""`: Will map to `"geo"` if you use `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
Available values: "off", "geo", "random", "dynamic_latency", "proximity", "least_outstanding_requests", "least_connections", "".



<a id="nestedatt--random_steering"></a>
### Nested Schema for `random_steering`

Optional:

- `default_weight` (Number) The default weight for pools in the load balancer that are not specified in the pool_weights map.
- `pool_weights` (Map of Number) A mapping of pool IDs to custom weights. The weight is relative to other pools in the load balancer.


<a id="nestedatt--rules"></a>
### Nested Schema for `rules`

Optional:

- `condition` (String) The condition expressions to evaluate. If the condition evaluates to true, the overrides or fixed_response in this rule will be applied. An empty condition is always true. For more details on condition expressions, please see https://developers.cloudflare.com/load-balancing/understand-basics/load-balancing-rules/expressions.
- `disabled` (Boolean) Disable this specific rule. It will no longer be evaluated by this load balancer.
- `fixed_response` (Attributes) A collection of fields used to directly respond to the eyeball instead of routing to a pool. If a fixed_response is supplied the rule will be marked as terminates. (see [below for nested schema](#nestedatt--rules--fixed_response))
- `name` (String) Name of this rule. Only used for human readability.
- `overrides` (Attributes) A collection of overrides to apply to the load balancer when this rule's condition is true. All fields are optional. (see [below for nested schema](#nestedatt--rules--overrides))
- `priority` (Number) The order in which rules should be executed in relation to each other. Lower values are executed first. Values do not need to be sequential. If all priorities are omitted or zero, the API assigns priorities based on the array order of the rules field.
- `terminates` (Boolean) If this rule's condition is true, this causes rule evaluation to stop after processing this rule.

<a id="nestedatt--rules--fixed_response"></a>
### Nested Schema for `rules.fixed_response`

Optional:

- `content_type` (String) The http 'Content-Type' header to include in the response.
- `location` (String) The http 'Location' header to include in the response.
- `message_body` (String) Text to include as the http body.
- `status_code` (Number) The http status code to respond with.


<a id="nestedatt--rules--overrides"></a>
### Nested Schema for `rules.overrides`

Optional:

- `adaptive_routing` (Attributes) Controls features that modify the routing of requests to pools and origins in response to dynamic conditions, such as during the interval between active health monitoring requests. For example, zero-downtime failover occurs immediately when an origin becomes unavailable due to HTTP 521, 522, or 523 response codes. If there is another healthy origin in the same pool, the request is retried once against this alternate origin. (see [below for nested schema](#nestedatt--rules--overrides--adaptive_routing))
- `country_pools` (Map of List of String) A mapping of country codes to a list of pool IDs (ordered by their failover priority) for the given country. Any country not explicitly defined will fall back to using the corresponding region_pool mapping if it exists else to default_pools.
- `default_pools` (List of String) A list of pool IDs ordered by their failover priority. Pools defined here are used by default, or when region_pools are not configured for a given region.
- `fallback_pool` (String) The pool ID to use when all other pools are detected as unhealthy.
- `location_strategy` (Attributes) Controls location-based steering for non-proxied requests. See `steering_policy` to learn how steering is affected. (see [below for nested schema](#nestedatt--rules--overrides--location_strategy))
- `pop_pools` (Map of List of String) Enterprise only: A mapping of Cloudflare PoP identifiers to a list of pool IDs (ordered by their failover priority) for the PoP (datacenter). Any PoPs not explicitly defined will fall back to using the corresponding country_pool, then region_pool mapping if it exists else to default_pools.
- `random_steering` (Attributes) Configures pool weights for a rule override. An omitted default_weight inherits the load balancer's default weight.
- `steering_policy="random"`: A random pool is selected with probability proportional to pool weights.
- `steering_policy="least_outstanding_requests"`: Use pool weights to scale each pool's outstanding requests.
- `steering_policy="least_connections"`: Use pool weights to scale each pool's open connections. (see [below for nested schema](#nestedatt--rules--overrides--random_steering))
- `region_pools` (Map of List of String) A mapping of region codes to a list of pool IDs (ordered by their failover priority) for the given region. Any regions not explicitly defined will fall back to using default_pools.
- `session_affinity` (String) Specifies the type of session affinity the load balancer should use unless specified as `"none"`. The supported types are: - `"cookie"`: On the first request to a proxied load balancer, a cookie is generated, encoding information of which origin the request will be forwarded to. Subsequent requests, by the same client to the same load balancer, will be sent to the origin server the cookie encodes, for the duration of the cookie and as long as the origin server remains healthy. If the cookie has expired or the origin server is unhealthy, then a new origin server is calculated and used. - `"ip_cookie"`: Behaves the same as `"cookie"` except the initial origin selection is stable and based on the client's ip address. - `"header"`: On the first request to a proxied load balancer, a session key based on the configured HTTP headers (see `session_affinity_attributes.headers`) is generated, encoding the request headers used for storing in the load balancer session state which origin the request will be forwarded to. Subsequent requests to the load balancer with the same headers will be sent to the same origin server, for the duration of the session and as long as the origin server remains healthy. If the session has been idle for the duration of `session_affinity_ttl` seconds or the origin server is unhealthy, then a new origin server is calculated and used. See `headers` in `session_affinity_attributes` for additional required configuration.
Available values: "none", "cookie", "ip_cookie", "header".
- `session_affinity_attributes` (Attributes) Configures attributes for session affinity. (see [below for nested schema](#nestedatt--rules--overrides--session_affinity_attributes))
- `session_affinity_ttl` (Number) Time, in seconds, until a client's session expires after being created. Once the expiry time has been reached, subsequent requests may get sent to a different origin server. The accepted ranges per `session_affinity` policy are: - `"cookie"` / `"ip_cookie"`: The current default of 23 hours will be used unless explicitly set. The accepted range of values is between [1800, 604800]. - `"header"`: The current default of 1800 seconds will be used unless explicitly set. The accepted range of values is between [30, 3600]. Note: With session affinity by header, sessions only expire after they haven't been used for the number of seconds specified.
- `steering_policy` (String) Steering Policy for this load balancer.
- `"off"`: Use `default_pools`.
- `"geo"`: Use `region_pools`/`country_pools`/`pop_pools`. For non-proxied requests, the country for `country_pools` is determined by `location_strategy`.
- `"random"`: Select a pool randomly.
- `"dynamic_latency"`: Use round trip time to select the closest pool in default_pools (requires pool health checks).
- `"proximity"`: Use the pools' latitude and longitude to select the closest pool using the Cloudflare PoP location for proxied requests or the location determined by `location_strategy` for non-proxied requests.
- `"least_outstanding_requests"`: Select a pool by taking into consideration `random_steering` weights, as well as each pool's number of outstanding requests. Pools with more pending requests are weighted proportionately less relative to others.
- `"least_connections"`: Select a pool by taking into consideration `random_steering` weights, as well as each pool's number of open connections. Pools with more open connections are weighted proportionately less relative to others. Supported for HTTP/1 and HTTP/2 connections.
- `""`: Will map to `"geo"` if you use `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
Available values: "off", "geo", "random", "dynamic_latency", "proximity", "least_outstanding_requests", "least_connections", "".
- `ttl` (Number) Time to live (TTL) of the DNS entry for the IP address returned by this load balancer. This only applies to gray-clouded (unproxied) load balancers.

<a id="nestedatt--rules--overrides--adaptive_routing"></a>
### Nested Schema for `rules.overrides.adaptive_routing`

Optional:

- `failover_across_pools` (Boolean) Extends zero-downtime failover of requests to healthy origins from alternate pools, when no healthy alternate exists in the same pool, according to the failover order defined by traffic and origin steering. When set false (the default) zero-downtime failover will only occur between origins within the same pool. See `session_affinity_attributes` for control over when sessions are broken or reassigned.


<a id="nestedatt--rules--overrides--location_strategy"></a>
### Nested Schema for `rules.overrides.location_strategy`

Optional:

- `mode` (String) Determines the authoritative location when ECS is not preferred, does not exist in the request, or its GeoIP lookup is unsuccessful.
- `"pop"`: Use the Cloudflare PoP location.
- `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is unsuccessful, use the Cloudflare PoP location.
Available values: "pop", "resolver_ip".
- `prefer_ecs` (String) Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the authoritative location.
- `"always"`: Always prefer ECS.
- `"never"`: Never prefer ECS.
- `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
- `"geo"`: Prefer ECS only when `steering_policy="geo"`.
Available values: "always", "never", "proximity", "geo".


<a id="nestedatt--rules--overrides--random_steering"></a>
### Nested Schema for `rules.overrides.random_steering`

Optional:

- `default_weight` (Number) The default weight for pools in the load balancer that are not specified in the pool_weights map. An omitted value inherits the load balancer's default weight.
- `pool_weights` (Map of Number) A mapping of pool IDs to custom weights. The weight is relative to other pools in the load balancer.


<a id="nestedatt--rules--overrides--session_affinity_attributes"></a>
### Nested Schema for `rules.overrides.session_affinity_attributes`

Optional:

- `drain_duration` (Number) Configures the drain duration in seconds. This field is only used when session affinity is enabled on the load balancer.
- `headers` (List of String) Configures the names of HTTP headers to base session affinity on when header `session_affinity` is enabled. At least one HTTP header name must be provided. To specify the exact cookies to be used, include an item in the following format: `"cookie:<cookie-name-1>,<cookie-name-2>"` (example) where everything after the colon is a comma-separated list of cookie names. Providing only `"cookie"` will result in all cookies being used. The default max number of HTTP header names that can be provided depends on your plan: 5 for Enterprise, 1 for all other plans.
- `require_all_headers` (Boolean) When header `session_affinity` is enabled, this option can be used to specify how HTTP headers on load balancing requests will be used. The supported values are: - `"true"`: Load balancing requests must contain *all* of the HTTP headers specified by the `headers` session affinity attribute, otherwise sessions aren't created. - `"false"`: Load balancing requests must contain *at least one* of the HTTP headers specified by the `headers` session affinity attribute, otherwise sessions aren't created.
- `samesite` (String) Configures the SameSite attribute on session affinity cookie. Value "Auto" will be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note: when using value "None", the secure attribute can not be set to "Never".
Available values: "Auto", "Lax", "None", "Strict".
- `secure` (String) Configures the Secure attribute on session affinity cookie. Value "Always" indicates the Secure attribute will be set in the Set-Cookie header, "Never" indicates the Secure attribute will not be set, and "Auto" will set the Secure attribute depending if Always Use HTTPS is enabled.
Available values: "Auto", "Always", "Never".
- `zero_downtime_failover` (String) Configures the zero-downtime failover between origins within a pool when session affinity is enabled. This feature is currently incompatible with Argo, Tiered Cache, and Bandwidth Alliance. The supported values are: - `"none"`: No failover takes place for sessions pinned to the origin (default). - `"temporary"`: Traffic will be sent to another other healthy origin until the originally pinned origin is available; note that this can potentially result in heavy origin flapping. - `"sticky"`: The session affinity cookie is updated and subsequent requests are sent to the new origin. Note: Zero-downtime failover with sticky sessions is currently not supported for session affinity by header.
Available values: "none", "temporary", "sticky".




<a id="nestedatt--session_affinity_attributes"></a>
### Nested Schema for `session_affinity_attributes`

Optional:

- `drain_duration` (Number) Configures the drain duration in seconds. This field is only used when session affinity is enabled on the load balancer.
- `headers` (List of String) Configures the names of HTTP headers to base session affinity on when header `session_affinity` is enabled. At least one HTTP header name must be provided. To specify the exact cookies to be used, include an item in the following format: `"cookie:<cookie-name-1>,<cookie-name-2>"` (example) where everything after the colon is a comma-separated list of cookie names. Providing only `"cookie"` will result in all cookies being used. The default max number of HTTP header names that can be provided depends on your plan: 5 for Enterprise, 1 for all other plans.
- `require_all_headers` (Boolean) When header `session_affinity` is enabled, this option can be used to specify how HTTP headers on load balancing requests will be used. The supported values are: - `"true"`: Load balancing requests must contain *all* of the HTTP headers specified by the `headers` session affinity attribute, otherwise sessions aren't created. - `"false"`: Load balancing requests must contain *at least one* of the HTTP headers specified by the `headers` session affinity attribute, otherwise sessions aren't created.
- `samesite` (String) Configures the SameSite attribute on session affinity cookie. Value "Auto" will be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note: when using value "None", the secure attribute can not be set to "Never".
Available values: "Auto", "Lax", "None", "Strict".
- `secure` (String) Configures the Secure attribute on session affinity cookie. Value "Always" indicates the Secure attribute will be set in the Set-Cookie header, "Never" indicates the Secure attribute will not be set, and "Auto" will set the Secure attribute depending if Always Use HTTPS is enabled.
Available values: "Auto", "Always", "Never".
- `zero_downtime_failover` (String) Configures the zero-downtime failover between origins within a pool when session affinity is enabled. This feature is currently incompatible with Argo, Tiered Cache, and Bandwidth Alliance. The supported values are: - `"none"`: No failover takes place for sessions pinned to the origin (default). - `"temporary"`: Traffic will be sent to another other healthy origin until the originally pinned origin is available; note that this can potentially result in heavy origin flapping. - `"sticky"`: The session affinity cookie is updated and subsequent requests are sent to the new origin. Note: Zero-downtime failover with sticky sessions is currently not supported for session affinity by header.
Available values: "none", "temporary", "sticky".

## Import

Import is supported using the following syntax:

```shell
$ terraform import cloudflare_load_balancer.example '<zone_id>/<load_balancer_id>'
```


