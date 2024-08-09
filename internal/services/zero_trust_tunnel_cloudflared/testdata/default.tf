
resource "cloudflare_zero_trust_tunnel_cloudflared" "%[2]s" {
	account_id = "%[1]s"
	name       = "%[2]s"
	secret     = "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="
}

data "cloudflare_zero_trust_tunnel_cloudflared" "%[2]s" {
	account_id = "%[1]s"
	name       = "%[2]s"
	is_deleted = false
	depends_on = [cloudflare_zero_trust_tunnel_cloudflared.%[2]s]
}
