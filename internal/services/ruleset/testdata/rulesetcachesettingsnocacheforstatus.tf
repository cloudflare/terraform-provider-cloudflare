
resource "cloudflare_ruleset" "%[1]s" {
  zone_id     = "%[2]s"
  name        = "%[1]s"
  description = "%[1]s ruleset description"
  kind        = "zone"
  phase       = "http_request_cache_settings"

  rules = [{
    action = "set_cache_settings"
    action_parameters = {
      edge_ttl = {
        mode    = "override_origin"
        default = 60 * 60 * 24 * 30 // 30 days
        status_code_ttl = [{
          status_code_range = { 
    from = 400
            to   = 500 
}
          value = 0
        }]
      }
    }
    expression  = "true"
    description = "%[1]s set cache settings rule"
    enabled     = true
  }]
}