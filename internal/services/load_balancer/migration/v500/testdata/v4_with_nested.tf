resource "cloudflare_load_balancer" "%s" {
  zone_id              = "%s"
  name                 = "%s"
  fallback_pool_id     = "%s"
  default_pool_ids     = ["%s"]
  steering_policy      = "off"
  ttl                  = 30
  session_affinity     = "cookie"
  session_affinity_ttl = 1800

  session_affinity_attributes {
    samesite              = "Lax"
    secure                = "Always"
    drain_duration        = 100
    zero_downtime_failover = "sticky"
  }

  adaptive_routing {
    failover_across_pools = true
  }

  location_strategy {
    prefer_ecs = "proximity"
    mode       = "pop"
  }

  random_steering {
    default_weight = 0.5
  }
}
