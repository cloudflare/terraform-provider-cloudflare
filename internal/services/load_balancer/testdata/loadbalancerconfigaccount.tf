resource "cloudflare_load_balancer" "%[3]s" {
  # Account-level LB (no zone_id)
  name    = "tf-testacc-lb-account-%[3]s"

  default_pools = ["${cloudflare_load_balancer_pool.%[1]s.id}"]
  fallback_pool = "${cloudflare_load_balancer_pool.%[1]s.id}"

  region_pools = {
    WNAM = ["${cloudflare_load_balancer_pool.%[1]s.id}"]
    SSAM = ["${cloudflare_load_balancer_pool.%[1]s.id}"]
  }
  pop_pools = {
    LHR = ["${cloudflare_load_balancer_pool.%[1]s.id}"]
  }
  steering_policy  = "geo"
  proxied          = true
  random_steering = {
    default_weight = 1.00
  }
}
