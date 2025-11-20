resource "cloudflare_zero_trust_gateway_policy" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[1]s"
  description = "Internal DNS resolve policy"
  action      = "resolve"
  filters     = ["dns_resolver"]
  traffic     = "any(dns.domains[*] == \"internal.example.com\")"

  rule_settings = {
    resolve_dns_internally = {
      view_id  = "internal-view-123"
      fallback = "public_dns"
    }
  }
}