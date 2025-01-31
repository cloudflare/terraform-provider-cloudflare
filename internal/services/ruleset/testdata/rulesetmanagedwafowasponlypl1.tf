
resource "cloudflare_ruleset" "%[1]s" {
  zone_id     = "%[3]s"
  name        = "%[2]s"
  description = "%[1]s ruleset description"
  kind        = "zone"
  phase       = "http_request_firewall_managed"

  # disable PL2, PL3 and PL4
  rules = [{
    action = "execute"
    action_parameters = {
      id = "4814384a9e5d4991b9815dcfc25d2f1f"
      overrides = { categories = [{
        category = "paranoia-level-2"
        enabled  = false
        },
        {
          category = "paranoia-level-3"
          enabled  = false
        },
        {
          category = "paranoia-level-4"
          enabled  = false
        }]



        rules = [{
          id              = "6179ae15870a4bb7b2d480d4843b323c"
          action          = "block"
          score_threshold = 60
          enabled         = true
      }] }
    }
    expression  = "true"
    description = "zone"
    enabled     = true
  }]
}