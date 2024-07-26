
resource "cloudflare_teams_rule" "%[1]s" {
  name = "%[1]s"
  account_id = "%[2]s"
  description = "desc"
  precedence = 12302
  action = "block"
  filters = ["dns"]
  traffic = "any(dns.domains[*] == \"example.com\")"
  rule_settings = {
  block_page_enabled = true
    block_page_reason = "cuz"
    insecure_disable_dnssec_validation = false
	egress = {
    ipv4 = "203.0.113.1"
		ipv6 = "2001:db8::/32"
  }
	untrusted_cert = {
    action = "error"
  }
	payload_log = {
    enabled = true
  }
}
}
