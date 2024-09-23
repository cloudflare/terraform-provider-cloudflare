
resource "cloudflare_ruleset" "%[1]s" {
  zone_id     = "%[2]s"
  name        = "%[1]s"
  description = "set cache settings for the request"
  kind        = "zone"
  phase       = "http_request_cache_settings"
  rules = [{
    action      = "set_cache_settings"
    description = "example"
    enabled     = true
    expression  = "(http.host eq \"example.com\" and starts_with(http.request.uri.path, \"/example\"))"
    action_parameters = {
      cache = true
      browser_ttl = {
        mode    = "bypass_by_default"
        default = 100
      }
    }
  }]
}
