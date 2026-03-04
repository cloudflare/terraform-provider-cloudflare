variable "zone_id" {}

resource "cloudflare_ruleset" "my_ruleset" {
  zone_id = var.zone_id
  name    = "My ruleset"
  phase   = "http_response_cache_settings"
  kind    = "zone"
  rules = [
    {
      expression = "any(http.request.headers.values[*] == \"application/json\")"
      action     = "set_cache_control"
      action_parameters = {
        max_age = {
          operation = "remove"
        }
        stale_while_revalidate = {
          operation = "set"
          value     = 12345
        }
        no_cache = {
          operation = "set"
        }
        public = {
          operation = "remove"
        }
        immutable = {
          operation = "set"
        }
        must_revalidate = {
          operation = "set"
        }
        proxy_revalidate = {
          operation = "remove"
        }
        no_transform = {
          operation = "remove"
        }
      }
    }
  ]
}

data "cloudflare_ruleset" "my_ruleset" {
  zone_id = var.zone_id
  id      = cloudflare_ruleset.my_ruleset.id
}
