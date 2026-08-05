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
  networks = ["string"]
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
