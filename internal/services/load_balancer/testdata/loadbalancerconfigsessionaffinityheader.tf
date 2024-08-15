resource "cloudflare_load_balancer" "%[3]s" {
  zone_id = "%[1]s"
  name = "tf-testacc-lb-session-affinity-%[3]s.%[2]s"
  fallback_pool_id = "${cloudflare_load_balancer_pool.%[3]s.id}"
  default_pool_ids = ["${cloudflare_load_balancer_pool.%[3]s.id}"]
  session_affinity = "header"
  session_affinity_ttl = 1800
  session_affinity_attributes = {
    samesite = "Auto"
    secure = "Auto"
    drain_duration = 60
    zero_downtime_failover = "temporary"
	headers = ["x-custom"]
	require_all_headers = true
  }
}
