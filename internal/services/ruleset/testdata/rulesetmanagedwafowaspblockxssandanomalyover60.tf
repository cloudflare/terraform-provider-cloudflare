
resource "cloudflare_ruleset" "%[1]s" {
  zone_id     = "%[3]s"
  name        = "%[2]s"
  description = "%[1]s ruleset description"
  kind        = "zone"
  phase       = "http_request_firewall_managed"

  # enable all "XSS" rules
  rules = [{
    action = "execute"
    action_parameters = {
      id = "efb7b8c949ac4650a09736fc376e9aee"
      overrides = { categories = [{
        category = "xss"
        action   = "block"
        enabled  = true
      }] }
    }
    expression  = "true"
    description = "zone"
    enabled     = true
    },
    {
      action = "execute"
      action_parameters = {
        id = "4814384a9e5d4991b9815dcfc25d2f1f"
        overrides = { rules = [{
          id              = "6179ae15870a4bb7b2d480d4843b323c"
          action          = "block"
          score_threshold = 60
        }] }
      }
      expression  = "true"
      description = "zone"
      enabled     = true
  }]

  # set Anomaly Score for 60+ (low)
}