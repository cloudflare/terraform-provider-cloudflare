resource "cloudflare_load_balancer" "%[3]s" {
  zone_id = "%[1]s"
  name = "tf-testacc-lb-random-steering-%[3]s.%[2]s"
  fallback_pool = "${cloudflare_load_balancer_pool.%[3]s.id}"
  default_pools = ["${cloudflare_load_balancer_pool.%[3]s.id}"]
  random_steering = {
    pool_weights = {
      "${cloudflare_load_balancer_pool.%[3]s.id}" = 0.3
    }
    default_weight = 0.9
  }
  session_affinity = "none"
}
