resource "cloudflare_zero_trust_gateway_policy" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[1]s"
  description = "DNS override policy"
  precedence  = 12400
  action      = "override"
  filters     = ["dns"]
  traffic     = "any(dns.domains[*] == \"example.com\")"

  rule_settings = {
    override_ips  = ["192.0.2.1", "192.0.2.2"]
  }
}