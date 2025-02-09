
resource "cloudflare_ruleset" "%[1]s" {
  zone_id     = "%[3]s"
  name        = "%[2]s"
  description = "%[1]s ruleset description"
  kind        = "zone"
  phase       = "http_request_cache_settings"

  rules = [{
    action = "set_cache_settings"
    action_parameters = {
      edge_ttl = {
        mode    = "override_origin"
        default = 60
      }
      browser_ttl = {
        mode = "respect_origin"
      }
    }
    expression  = "true"
    description = "%[1]s set cache settings rule"
    enabled     = true
  }]
}