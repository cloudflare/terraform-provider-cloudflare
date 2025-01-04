
resource "cloudflare_zero_trust_gateway_policy" "%[1]s" {
  name = "%[1]s"
  account_id = "%[2]s"
  description = "desc"
  precedence = 12302
  action = "block"
  filters = ["dns"]
  traffic = "any(dns.domains[*] == \"example.com\")"
}
