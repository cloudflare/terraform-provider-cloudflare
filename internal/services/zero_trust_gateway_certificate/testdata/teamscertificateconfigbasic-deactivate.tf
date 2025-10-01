
resource "cloudflare_zero_trust_gateway_certificate" "%[1]s" {
	account_id  = "%[2]s"
	validity_period_days = 1000 // 3 years
	activate = false
}
