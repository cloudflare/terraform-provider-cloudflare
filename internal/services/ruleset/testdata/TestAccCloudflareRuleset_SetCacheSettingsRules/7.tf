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
          cache_by_device_type = true
          custom_key = {
            query_string = {
              exclude = {
                list = ["foo"]
              }
            }
          }
        }
      }
    }
  ]
}
