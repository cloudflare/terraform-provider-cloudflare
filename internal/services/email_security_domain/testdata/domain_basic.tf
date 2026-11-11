resource "cloudflare_email_security_domain" "%[1]s" {
  account_id             = "%[2]s"
  domain                 = "%[1]s.xzy"
  allowed_delivery_modes = ["DIRECT"]
  drop_dispositions      = []
  ip_restrictions        = []
  regions                = ["US"]
  transport              = "transport.%[1]s.xzy"
  lookback_hops          = 1
  require_tls_inbound    = false
  require_tls_outbound   = false
}
