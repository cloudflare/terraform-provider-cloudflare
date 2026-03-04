resource "cloudflare_load_balancer" "%s" {
  zone_id           = "%s"
  name              = "%s"
  fallback_pool_id  = "%s"
  default_pool_ids  = ["%s"]
  steering_policy   = "geo"
  ttl               = 30

  region_pools {
    region   = "WNAM"
    pool_ids = ["%s"]
  }

  pop_pools {
    pop      = "LAX"
    pool_ids = ["%s"]
  }

  country_pools {
    country  = "US"
    pool_ids = ["%s"]
  }
}
