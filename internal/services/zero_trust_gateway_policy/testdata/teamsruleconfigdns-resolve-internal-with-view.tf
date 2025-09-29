resource "cloudflare_account_dns_settings_internal_view" "%[1]s_view" {
  account_id = "%[2]s"
  name       = "%[1]s-view"
  zones      = []  # Create view without zones first
}

resource "cloudflare_zero_trust_gateway_policy" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[1]s"
  description = "Internal DNS resolve policy with view"
  action      = "resolve"
  filters     = ["dns_resolver"]
  traffic     = "any(dns.domains[*] == \"internal.example.com\")"

  rule_settings = {
    resolve_dns_internally = {
      view_id  = cloudflare_account_dns_settings_internal_view.%[1]s_view.id
      fallback = "public_dns"
    }
  }
}