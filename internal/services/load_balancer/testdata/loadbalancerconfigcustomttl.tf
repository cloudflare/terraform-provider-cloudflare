resource "cloudflare_load_balancer" "%[3]s" {
  zone_id = "%[1]s"
  name    = "tf-testacc-lb-custom-ttl-%[3]s.%[2]s"

  default_pools = ["${cloudflare_load_balancer_pool.%[3]s.id}"]
  fallback_pool = "${cloudflare_load_balancer_pool.%[3]s.id}"

  region_pools = {
    WNAM = ["${cloudflare_load_balancer_pool.%[3]s.id}"]
    SSAM = ["${cloudflare_load_balancer_pool.%[3]s.id}"]
  }
  pop_pools = {
    LHR = ["${cloudflare_load_balancer_pool.%[3]s.id}"]
  }
  session_affinity     = "cookie"
  session_affinity_ttl = 5000
  steering_policy      = "geo"
  proxied              = true
  random_steering = {
    default_weight = 1.00
  }
}
