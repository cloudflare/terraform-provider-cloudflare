resource "cloudflare_load_balancer" "%[3]s" {
  zone_id = "%[1]s"
  name    = "tf-testacc-lb-custom-port-%[3]s.%[2]s"

  default_pools = ["${cloudflare_load_balancer_pool.%[3]s.id}"]
  fallback_pool = "${cloudflare_load_balancer_pool.%[3]s.id}"

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
