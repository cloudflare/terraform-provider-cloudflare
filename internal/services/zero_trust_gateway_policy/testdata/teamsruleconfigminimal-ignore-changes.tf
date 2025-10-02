resource "cloudflare_zero_trust_gateway_policy" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  action     = "block"
  filters    = ["dns"]
  traffic    = "any(dns.domains[*] == \"minimal.example.com\")"
  
  lifecycle {
    ignore_changes = [
      rule_settings,
      precedence,
      created_at,
      updated_at,
      version,
      deleted_at,
      read_only,
      sharable,
      source_account,
      warning_status,
      expiration,
      schedule
    ]
  }
}