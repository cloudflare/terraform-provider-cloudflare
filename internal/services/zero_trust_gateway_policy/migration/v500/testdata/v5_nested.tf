resource "cloudflare_zero_trust_gateway_policy" "%[1]s" {
  account_id  = "%[2]s"
  name        = "tf-test-nested-%[1]s"
  description = "Policy with nested blocks"
  precedence  = 10000
  action      = "block"
  enabled     = true
  filters     = ["dns"]
  traffic     = "any(dns.domains[*] in {\"blocked.com\"})"

  rule_settings = {
    block_page_enabled = true
    add_headers        = {}
    override_ips       = []

    notification_settings = {
      enabled     = true
      msg         = "Connection blocked"
      support_url = "https://support.example.com/"
    }
  }
}
