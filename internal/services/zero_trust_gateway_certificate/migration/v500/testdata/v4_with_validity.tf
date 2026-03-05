resource "cloudflare_zero_trust_gateway_certificate" "%s" {
  account_id           = "%s"
  gateway_managed      = true
  validity_period_days = 3650
}
