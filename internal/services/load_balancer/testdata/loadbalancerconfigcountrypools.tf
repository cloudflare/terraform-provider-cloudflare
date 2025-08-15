resource "cloudflare_load_balancer" "%[3]s" {
  zone_id = "%[1]s"
  name    = "tf-testacc-lb-country-pools-%[3]s.%[2]s"
  
  fallback_pool = "${cloudflare_load_balancer_pool.%[3]s.id}"
  default_pools = ["${cloudflare_load_balancer_pool.%[3]s.id}"]

  region_pools = {
    WNAM = ["${cloudflare_load_balancer_pool.%[3]s.id}"]
    IN   = ["${cloudflare_load_balancer_pool.%[3]s.id}"]
  }
  country_pools = {
    IN = ["${cloudflare_load_balancer_pool.%[3]s.id}"]
  }
  pop_pools = {
    LHR = ["{cloudflare_load_balancer_pool.%[3]s.id}"]
  }
  session_affinity = "cookie"
  steering_policy  = "geo"
  proxied          = true
  session_affinity_attributes = {
    samesite                = "Auto"
    secure                  = "Auto"
    drain_duration         = 0
    zero_downtime_failover = "none"
  }
  random_steering = {
    default_weight = 1.00
  }
}
