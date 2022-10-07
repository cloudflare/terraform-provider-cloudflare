# Example Usage
#
# Define a load balancer which always points to a pool we define below.
# In normal usage, would have different pools set for different pops
# (cloudflare points-of-presence) and/or for different regions.
# Within each pop or region we can define multiple pools in failover order.
resource "cloudflare_load_balancer" "bar" {
  zone_id          = "0da42c8d2132a9ddaf714f9e7c920711"
  name             = "example-load-balancer.example.com"
  fallback_pool_id = cloudflare_load_balancer_pool.foo.id
  default_pool_ids = [cloudflare_load_balancer_pool.foo.id]
  description      = "example load balancer using geo-balancing"
  proxied          = true
  steering_policy  = "geo"

  pop_pools {
    pop      = "LAX"
    pool_ids = [cloudflare_load_balancer_pool.foo.id]
  }

  country_pools {
    country  = "US"
    pool_ids = [cloudflare_load_balancer_pool.foo.id]
  }

  region_pools {
    region   = "WNAM"
    pool_ids = [cloudflare_load_balancer_pool.foo.id]
  }

  rules {
    name      = "example rule"
    condition = "http.request.uri.path contains \"testing\""
    fixed_response {
      message_body = "hello"
      status_code  = 200
      content_type = "html"
      location     = "www.example.com"
    }
  }
}

resource "cloudflare_load_balancer_pool" "foo" {
  name = "example-lb-pool"
  origins {
    name    = "example-1"
    address = "192.0.2.1"
    enabled = false
  }
}
