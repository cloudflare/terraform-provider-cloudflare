
resource "cloudflare_load_balancer" "%[3]s" {
  zone_id = "%[1]s"
  name = "tf-testacc-lb-random-steering-%[3]s.%[2]s"
  fallback_pool_id = "${cloudflare_load_balancer_pool.%[3]s.id}"
  default_pool_ids = ["${cloudflare_load_balancer_pool.%[3]s.id}"]
  random_steering = {
  pool_weights = {
      "${cloudflare_load_balancer_pool.%[3]s.id}" = 0.4
    }
    default_weight = 0.8
}
}
