resource "cloudflare_load_balancer" "%[3]s" {
  zone_id          = "%[1]s"
  name             = "tf-testacc-lb-%[3]s.%[2]s"
  steering_policy  = ""
  description      = "rules lb"
  fallback_pool = cloudflare_load_balancer_pool.%[3]s.id
  default_pools = ["${cloudflare_load_balancer_pool.%[3]s.id}"]
  rules = [
    {
      name      = "test rule 1"
      condition = "dns.qry.type == 28"
      overrides = { 
        steering_policy = "geo"
        session_affinity_attributes = {
          samesite               = "Auto"
          secure                 = "Auto"
          zero_downtime_failover = "sticky"
        }
        adaptive_routing = {
          failover_across_pools = true
        }
        location_strategy = {
          prefer_ecs = "always"
          mode       = "resolver_ip"
        }
        random_steering = {
          pool_weights = {
            "${cloudflare_load_balancer_pool.%[3]s.id}" = 0.4
          }
          default_weight = 0.2
        } 
      }
    },
    {
      name      = "test rule 2"
      condition = "dns.qry.type == 28"
      fixed_response = {
        message_body = "hello"
        status_code  = 200
        content_type = "html"
        location     = "www.example.com"
      }
    },
    {
      name      = "test rule 3"
      condition = "dns.qry.type == 28"
      overrides = {
        region_pools = {
          "ENAM": ["${cloudflare_load_balancer_pool.%[3]s.id}"]
        }
      }
    }
  ]
  session_affinity = "none"
}
