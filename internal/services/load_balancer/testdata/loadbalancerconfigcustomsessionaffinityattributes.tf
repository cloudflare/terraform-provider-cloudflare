resource "cloudflare_load_balancer" "%[3]s" {
  zone_id = "%[1]s"
  name    = "tf-testacc-lb-custom-session-affinity-%[3]s.%[2]s"

  default_pools = ["${cloudflare_load_balancer_pool.%[3]s.id}"]
  fallback_pool = "${cloudflare_load_balancer_pool.%[3]s.id}"

  steering_policy = "off"
  proxied         = true

  session_affinity_attributes = {
    samesite                = "Lax"
    secure                  = "Always"
    drain_duration          = 60
    zero_downtime_failover  = "temporary"
  }
}
