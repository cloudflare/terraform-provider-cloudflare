
resource "cloudflare_ruleset" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[1]s managed WAF"
  description = "%[1]s managed WAF ruleset description"
  kind        = "root"
  phase       = "http_request_firewall_managed"

  rules = [{
    action = "execute"
    action_parameters = {
      id = "4814384a9e5d4991b9815dcfc25d2f1f"
      overrides = { rules = [{
        id              = "6179ae15870a4bb7b2d480d4843b323c"
        action          = "block"
        score_threshold = 25
        }]
      enabled = true }
      matched_data = {
        public_key = "zpUlcpNtaNiSUN6LL6NiNz8XgIJZWWG3iSZDdPbMszM="
      }
    }
    expression  = "(cf.zone.name eq \"%[3]s\") and (cf.zone.plan eq \"ENT\")"
    description = "Account OWASP %[3]s"
    enabled     = true
  }]
}
