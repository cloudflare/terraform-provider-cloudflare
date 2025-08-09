
resource "cloudflare_ruleset" "%[1]s_account_custom_firewall" {
  account_id  = "%[3]s"
  name        = "Custom Ruleset for my account"
  description = "example block rule"
  kind        = "custom"
  phase       = "http_request_firewall_custom"

  rules = [{
    action      = "block"
    expression  = "(http.host eq \"%[4]s\")"
    description = "SID"
    enabled     = true
  }]
}

resource "cloudflare_ruleset" "%[1]s_account_custom_firewall_root" {
  account_id  = "%[3]s"
  name        = "Firewall Custom root"
  description = ""
  kind        = "root"
  phase       = "http_request_firewall_custom"

  rules = [{
    action = "execute"
    action_parameters = {
      id = cloudflare_ruleset.%[1]s_account_custom_firewall.id
    }
    expression  = "(cf.zone.name eq \"example.com\") and (cf.zone.plan eq \"ENT\")"
    description = ""
    enabled     = true
  }]
}
