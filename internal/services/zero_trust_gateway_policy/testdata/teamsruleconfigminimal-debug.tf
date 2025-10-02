resource "cloudflare_zero_trust_gateway_policy" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  action     = "block"
  filters    = ["dns"]
  traffic    = "any(dns.domains[*] == \"minimal.example.com\")"
}