resource "cloudflare_load_balancer" "%s" {
  zone_id              = "%s"
  name                 = "%s"
  fallback_pool        = "%s"
  default_pools        = ["%s"]
  description          = "Load balancer for tf-migrate test"
  ttl                  = 30
  session_affinity     = "cookie"
  session_affinity_ttl = 1800
  enabled              = true
  steering_policy      = "dynamic_latency"
}
