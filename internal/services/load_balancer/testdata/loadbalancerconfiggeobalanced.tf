resource "cloudflare_load_balancer" "%[3]s" {
  zone_id = "%[1]s"
  name = "tf-testacc-lb-%[3]s.%[2]s"
  fallback_pool = "${cloudflare_load_balancer_pool.%[3]s.id}"
  default_pools = ["${cloudflare_load_balancer_pool.%[3]s.id}"]
  description = "tf-acctest load balancer using geo-balancing"
  proxied = true // can't set ttl with proxied
  steering_policy = "geo"
  pop_pools ={
    "LAX": ["${cloudflare_load_balancer_pool.%[3]s.id}"]
  }
  country_pools = {
    "US": ["${cloudflare_load_balancer_pool.%[3]s.id}"]
  }
  region_pools = {
    "WNAM": ["${cloudflare_load_balancer_pool.%[3]s.id}"]
  }
  session_affinity = "none"
}
