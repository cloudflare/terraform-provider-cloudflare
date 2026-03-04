variable "zone_id" {}

resource "cloudflare_ruleset" "my_ruleset" {
  zone_id = var.zone_id
  name    = "My cache tags ruleset"
  phase   = "http_response_cache_settings"
  kind    = "zone"
  rules = [
    {
      expression = "http.request.uri.path contains \"/content\""
      action     = "set_cache_tags"
      action_parameters = {
        operation = "set"
        values     = ["content", "public"]
      }
    }
  ]
}

data "cloudflare_ruleset" "my_ruleset" {
  zone_id = var.zone_id
  id      = cloudflare_ruleset.my_ruleset.id
}