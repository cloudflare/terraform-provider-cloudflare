resource "cloudflare_zero_trust_gateway_policy" "%[1]s" {
  account_id  = "%[2]s"
  name        = "tf-test-minimal-%[1]s"
  description = "Minimal policy for migration testing"
  precedence  = 10000
  action      = "block"
  filters     = ["dns"]
  traffic     = "any(dns.domains[*] in {\"example.com\"})"
}
