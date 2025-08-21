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
        browser_ttl = {
          mode    = "override_origin"
          default = 60
        }
        cache_key = {
          cache_by_device_type       = false
          cache_deception_armor      = false
          custom_key                 = {}
          ignore_query_strings_order = false
        }
        cache_reserve = {
          eligible          = true
          minimum_file_size = 1024
        }
        edge_ttl = {
          default = 60
          mode    = "override_origin"
          status_code_ttl = [
            {
              status_code_range = {
                from = "500"
              }
              value = -1
            },
            {
              status_code_range = {
                to = "199"
              }
              value = 0
            },
            {
              status_code_range = {
                from = "200"
                to   = "399"
              }
              value = 1
            },
            {
              status_code = 400
              value       = 2
            }
          ]
        }
        origin_cache_control       = true
        origin_error_page_passthru = true
        respect_strong_etags       = true
        serve_stale = {
          disable_stale_while_updating = false
        }
      }
    }
  ]
}
