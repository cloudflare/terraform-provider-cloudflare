resource "cloudflare_zero_trust_gateway_policy" "%[1]s" {
  account_id    = "%[2]s"
  name          = "%[1]s"
  description   = "Maximal policy with all options"
  action        = "block"
  enabled       = true
  precedence    = 12500
  filters       = ["dns"]
  traffic       = "any(dns.domains[*] == \"blocked.example.com\")"
  identity      = "any(identity.groups.name[*] in {\"finance\"})"
  device_posture = "any(device_posture.checks.passed[*] == \"device-check-id\")"

  rule_settings = {
    block_page_enabled     = true
    block_reason          = "Policy violation"
    ip_categories         = true
    ip_indicator_feeds    = true
  }
}