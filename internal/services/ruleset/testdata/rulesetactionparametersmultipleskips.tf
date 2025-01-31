
resource "cloudflare_ruleset" "%[1]s" {
  zone_id     = "%[3]s"
  name        = "%[2]s"
  description = "%[1]s ruleset description"
  kind        = "zone"
  phase       = "http_request_firewall_managed"

  rules = [{
    action = "skip"
    action_parameters = {
      rulesets = ["efb7b8c949ac4650a09736fc376e9aee"]
    }
    expression  = "(cf.zone.name eq \"domain.xyz\" and http.request.uri.query contains \"skip=rulesets\")"
    description = "skip Cloudflare Manage ruleset"
    enabled     = true
    logging = {
      enabled = true
    }
    },
    {
      action = "skip"
      action_parameters = {
        # efb7b8c949ac4650a09736fc376e9aee is the ruleset ID of the Cloudflare Managed rules
        rules = {
          "efb7b8c949ac4650a09736fc376e9aee" = ["5de7edfa648c4d6891dc3e7f84534ffa", "e3a567afc347477d9702d9047e97d760"]
        }
      }
      expression  = "(cf.zone.name eq \"domain.xyz\" and http.request.uri.query contains \"skip=rules\")"
      description = "skip Wordpress rule and SQLi rule"
      enabled     = true
      logging = {
        enabled = true
      }
    },
    {
      action = "execute"
      action_parameters = {
        id = "efb7b8c949ac4650a09736fc376e9aee"
        overrides = { rules = [{
          id      = "5de7edfa648c4d6891dc3e7f84534ffa"
          action  = "block"
          enabled = true
          },
          {
            id      = "75a0060762034a6cb663fd51a02344cb"
            action  = "log"
            enabled = true
          }]

          categories = [{
            category = "wordpress"
            action   = "js_challenge"
            enabled  = true
        }] }
      }
      expression  = "true"
      description = "Execute Cloudflare Managed Ruleset on my zone-level phase entry point ruleset"
      enabled     = true
  }]


}
