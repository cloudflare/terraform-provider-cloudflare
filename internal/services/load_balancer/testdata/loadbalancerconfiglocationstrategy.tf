resource "cloudflare_load_balancer" "%[3]s" {
  zone_id           = "%[1]s"
  name              = "tf-testacc-lb-location-strategy-%[3]s.%[2]s"
  fallback_pool     = "${cloudflare_load_balancer_pool.%[3]s.id}"
  default_pools     = ["${cloudflare_load_balancer_pool.%[3]s.id}"]
  location_strategy = {
    prefer_ecs = "proximity"
    mode = "pop"
  }
  session_affinity = "cookie"
}
