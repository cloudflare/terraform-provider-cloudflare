resource "cloudflare_load_balancer" "%[3]s" {
  zone_id          = "%[1]s"
  name             = "%[3]s.%[2]s"
  default_pools = ["${cloudflare_load_balancer_pool.%[1]s.id}"]
  fallback_pool = "${cloudflare_load_balancer_pool.%[1]s.id}"
  description      = "Load Balancer for %[2]s"
  ttl              = 30
  proxied          = false
  steering_policy  = "off"

  location_strategy {
    prefer_ecs = "always"
    mode       = "resolver_ip"
  }
}
