resource "cloudflare_load_balancer" "%[3]s" {
  zone_id          = "%[1]s"
  name             = "%[3]s.%[2]s"
  fallback_pool = "${cloudflare_load_balancer_pool.%[1]s.id}"
  default_pools = ["${cloudflare_load_balancer_pool.%[1]s.id}"]
  description      = "Load Balancer for %[2]s"
  ttl              = 30
  proxied          = false
  steering_policy  = "off"

  adaptive_routing {
    failover_across_pools = true
  }
}
