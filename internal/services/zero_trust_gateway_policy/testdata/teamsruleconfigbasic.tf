resource "cloudflare_zero_trust_gateway_policy" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[1]s"
  description = "Basic DNS policy for testing"
  precedence  = 12350
  action      = "block"
  filters     = ["dns"]
  traffic     = "any(dns.domains[*] == \"basic.example.com\")"

  rule_settings = {
    block_page_enabled = true
    block_reason       = "Basic test policy"
  }
}