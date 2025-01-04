
resource "cloudflare_load_balancer" "%[3]s" {
  zone_id = "%[1]s"
  name = "tf-testacc-lb-%[3]s.%[2]s"
  fallback_pool = "${cloudflare_load_balancer_pool.%[3]s.id}"
  default_pools = ["${cloudflare_load_balancer_pool.%[3]s.id}"]
  description = "tf-acctest load balancer using least connections steering"
  proxied = true
  steering_policy = "least_connections"
  rules =[{
    name = "test rule 1"
    condition = "dns.qry.type == 28"
    overrides = { steering_policy = "least_connections" }
  }]
  session_affinity = "cookie"
}
