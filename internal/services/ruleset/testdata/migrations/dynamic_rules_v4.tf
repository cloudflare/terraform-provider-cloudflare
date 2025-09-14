locals {
  rule_configs = [
    {
      expression = "ip.src eq 1.1.1.1"
      action     = "block"
      description = "Block specific IP"
    },
    {
      expression = "http.request.uri.path contains \"/admin\""
      action     = "challenge"
      description = "Challenge admin paths"
    },
    {
      expression = "ip.src eq 2.2.2.2"
      action     = "log"
      description = "Log requests from IP"
    }
  ]
}

resource "cloudflare_ruleset" "%[2]s" {
  zone_id = "%[1]s"
  name    = "My ruleset %[2]s"
  phase   = "http_request_firewall_custom"
  kind    = "zone"

  dynamic "rules" {
    for_each = local.rule_configs
    content {
      expression  = rules.value.expression
      action      = rules.value.action
      description = rules.value.description
      enabled     = true
    }
  }
}