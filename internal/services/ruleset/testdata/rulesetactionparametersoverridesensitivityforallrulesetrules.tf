
resource "cloudflare_ruleset" "%[1]s" {
  zone_id     = "%[3]s"
  name        = "%[2]s"
  description = "%[1]s ruleset description"
  kind        = "zone"
  phase       = "ddos_l7"

  rules = [{
    action = "execute"
    action_parameters = {
      id = "4d21379b4f9f4bb088e0729962c8b3cf"
      overrides = { action = "log"
      sensitivity_level = "low" }
    }
    expression  = "true"
    description = "override HTTP DDoS ruleset rule"
    enabled     = true
  }]
}