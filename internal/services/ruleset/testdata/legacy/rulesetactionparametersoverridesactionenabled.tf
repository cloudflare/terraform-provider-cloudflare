
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
      overrides = { action = "log"
      enabled = true }
    }
    expression  = "true"
    description = "Execute all rules in Cloudflare Managed Ruleset in log mode on my zone-level phase entry point ruleset"
    enabled     = true
  }]
}