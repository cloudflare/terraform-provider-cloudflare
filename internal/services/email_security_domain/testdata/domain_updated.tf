resource "cloudflare_email_security_domain" "%[1]s" {
  account_id             = "%[2]s"
  domain                 = "%[1]s.xzy"
  allowed_delivery_modes = ["DIRECT"]
  drop_dispositions      = ["MALICIOUS", "SUSPICIOUS"]
  ip_restrictions        = ["192.0.2.0/24", "2001:db8::/32"]
  regions                = ["US", "DE"]
  transport              = "updated.%[1]s.xzy"
  lookback_hops          = 5
  require_tls_inbound    = true
  require_tls_outbound   = true
}
