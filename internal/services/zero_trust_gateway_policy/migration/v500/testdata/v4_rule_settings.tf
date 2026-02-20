resource "cloudflare_teams_rule" "%[1]s" {
  account_id  = "%[2]s"
  name        = "tf-test-settings-%[1]s"
  description = "Policy with rule settings"
  precedence  = 10000
  action      = "block"
  enabled     = true
  filters     = ["dns"]
  traffic     = "any(dns.domains[*] in {\"badsite.com\"})"

  rule_settings {
    block_page_enabled = true
    block_page_reason  = "Access blocked by company policy"
    ip_categories      = true
    add_headers        = {}
    override_ips       = []
  }
}
