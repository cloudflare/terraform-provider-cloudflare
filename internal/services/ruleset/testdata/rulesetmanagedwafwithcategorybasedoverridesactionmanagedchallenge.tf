
resource "cloudflare_ruleset" "%[1]s" {
  zone_id     = "%[3]s"
  name        = "%[2]s"
  description = "%[1]s ruleset description"
  kind        = "zone"
  phase       = "http_request_firewall_managed"

  rules = [{
    action = "execute"
    action_parameters = {
      id = "efb7b8c949ac4650a09736fc376e9aee"
      overrides = { categories = [{
        category = "wordpress"
        action   = "managed_challenge"
        enabled  = true
        }]
        rules = [{
          id      = "e3a567afc347477d9702d9047e97d760"
          action  = "managed_challenge"
          enabled = true
      }] }
    }

    expression  = "true"
    description = "overrides to only enable wordpress rules to managed_challenge"
    enabled     = true
  }]
}