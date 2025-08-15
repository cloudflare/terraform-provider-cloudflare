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
            cookie = {
              check_presence = ["myCookie1"]
              include        = ["myCookie2"]
            }
            header = {
              check_presence = ["my-header-1"]
              contains = {
                "my-header" = ["my-header-value"]
              }
              exclude_origin = false
              include        = ["my-header-2"]
            }
            host = {
              resolved = false
            }
            query_string = {
              include = {
                list = ["foo"]
              }
            }
            user = {
              device_type = false
              geo         = false
              lang        = false
            }
          }
        }
      }
    }
  ]
}
