
resource "cloudflare_ruleset" "%[1]s" {
  account_id  = "%[3]s"
  name        = "%[2]s"
  description = "This ruleset includes a skip rule whose logging is disabled."
  kind        = "root"
  phase       = "http_request_firewall_managed"

  rules = [{
    action = "skip"
    action_parameters = {
      ruleset = "current"
    }
    expression  = "(cf.zone.plan eq \"ENT\")"
    enabled     = true
    description = "example disabled logging"
    logging = {
      enabled = false
    }
  }]
}
