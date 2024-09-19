resource "cloudflare_load_balancer" "%[3]s" {
  zone_id = "%[1]s"
  name = "tf-testacc-lb-%[3]s.%[2]s"
  fallback_pool_id = "${cloudflare_load_balancer_pool.%[3]s.id}"
  default_pool_ids = ["${cloudflare_load_balancer_pool.%[3]s.id}"]
  description = "tf-acctest load balancer using pop/country geo-balancing updated to region geo-balancing"
  proxied = true
  steering_policy = "geo"
  region_pools ={
    "WNAM": ["${cloudflare_load_balancer_pool.%[3]s.id}"]
  }
  session_affinity = "cookie"
}
