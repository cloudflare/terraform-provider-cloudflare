resource "cloudflare_load_balancer" "%s" {
  zone_id       = "%s"
  name          = "%s"
  fallback_pool = "%s"
  default_pools = ["%s"]

  region_pools = {
    "WNAM" = ["%s"]
  }

  pop_pools = {
    "LAX" = ["%s"]
  }

  country_pools = {
    "US" = ["%s"]
  }
}
