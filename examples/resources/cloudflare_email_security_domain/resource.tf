resource "cloudflare_email_security_domain" "example_email_security_domain" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  allowed_delivery_modes = ["DIRECT"]
  domain = "domain"
  drop_dispositions = ["MALICIOUS"]
  ip_restrictions = ["192.0.2.0/24", "2001:db8::/32"]
  regions = ["GLOBAL"]
  folder = "AllItems"
  integration_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
  lookback_hops = 1
  require_tls_inbound = true
  require_tls_outbound = true
  transport = "transport"
}
