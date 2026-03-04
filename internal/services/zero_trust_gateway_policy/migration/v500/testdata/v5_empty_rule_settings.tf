resource "cloudflare_zero_trust_gateway_policy" "%[1]s" {
  account_id  = "%[2]s"
  name        = "tf-test-empty-%[1]s"
  description = "Policy without rule settings"
  precedence  = 10000
  action      = "block"
  enabled     = false
  filters     = ["dns"]
  traffic     = "any(dns.domains[*] in {\"test.com\"})"
}
