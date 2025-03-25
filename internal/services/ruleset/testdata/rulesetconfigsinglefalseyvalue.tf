
resource "cloudflare_ruleset" "%[1]s" {
  zone_id     = "%[2]s"
  name        = "%[1]s"
  description = "%[1]s ruleset description"
  kind        = "zone"
  phase       = "http_config_settings"

  rules = [{
    action = "set_config"
    action_parameters = {
      bic = false
    }
    expression  = "true"
    description = "disable BIC"
    enabled     = true
  }]
}