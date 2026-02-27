variable "zone_id" {}

resource "cloudflare_ruleset" "my_ruleset" {
  zone_id = var.zone_id
  name    = "My cache tags ruleset"
  phase   = "http_response_cache_settings"
  kind    = "zone"
  rules = [
    {
      expression = "http.request.uri.path matches \"^/api/v[0-9]+/\""
      action     = "set_cache_tags"
      action_parameters = {
        operation  = "add"
        expression = "split(http.response.headers[\"cache-tag-ext\"][0], \",\", 1)"
      }
    }
  ]
}

data "cloudflare_ruleset" "my_ruleset" {
  zone_id = var.zone_id
  id      = cloudflare_ruleset.my_ruleset.id
}