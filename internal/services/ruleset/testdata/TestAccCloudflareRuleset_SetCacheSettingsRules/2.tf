variable "zone_id" {}

resource "cloudflare_ruleset" "my_ruleset" {
  zone_id = var.zone_id
  name    = "My ruleset"
  phase   = "http_request_cache_settings"
  kind    = "zone"
  rules = [
    {
      expression = "ip.src eq 1.1.1.1"
      action     = "set_cache_settings"
      action_parameters = {
        additional_cacheable_ports = ["8080"]
        browser_ttl = {
          mode = "respect_origin"
        }
        cache     = true
        cache_key = {}
        cache_reserve = {
          eligible = false
        }
        edge_ttl = {
          mode = "respect_origin"
        }
        origin_cache_control       = false
        origin_error_page_passthru = false
        read_timeout               = 900
        respect_strong_etags       = false
        serve_stale                = {}
      }
    }
  ]
}
