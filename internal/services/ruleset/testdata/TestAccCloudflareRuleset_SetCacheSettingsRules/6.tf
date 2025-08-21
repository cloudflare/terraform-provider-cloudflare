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
        cache_key = {
          custom_key = {
            header = {
              exclude_origin = true
            }
            host = {
              resolved = true
            }
            query_string = {
              include = {
                all = true
              }
            }
            user = {
              device_type = true
              geo         = true
              lang        = true
            }
          }
        }
      }
    }
  ]
}
