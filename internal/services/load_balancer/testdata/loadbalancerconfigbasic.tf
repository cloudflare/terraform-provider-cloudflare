resource "cloudflare_load_balancer" "%[3]s" {
  zone_id          = "%[1]s"
  name             = "tf-testacc-lb-%[3]s.%[2]s"
  steering_policy  = "off"
  session_affinity = "none"
  fallback_pool    = "${cloudflare_load_balancer_pool.%[3]s.id}"
  default_pools    = [
    "${cloudflare_load_balancer_pool.%[3]s.id}"
  ]
}
