resource "cloudflare_zero_trust_gateway_policy" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[1]s"
  description = "Minimal policy for testing"
  action      = "block"
  filters     = ["dns"]
  traffic     = "any(dns.domains[*] == \"minimal.example.com\")"
  
  # Include minimal rule_settings to reduce API drift
  rule_settings = {
    block_page_enabled = false
  }
}