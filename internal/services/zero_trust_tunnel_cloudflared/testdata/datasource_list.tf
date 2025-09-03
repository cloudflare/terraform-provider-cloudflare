resource "cloudflare_zero_trust_tunnel_cloudflared" "%[2]s" {
	account_id = "%[1]s"
	name       = "%[2]s"
	tunnel_secret     = "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="
}

data "cloudflare_zero_trust_tunnel_cloudflareds" "%[2]s" {
	account_id = "%[1]s"
	uuid = cloudflare_zero_trust_tunnel_cloudflared.%[2]s.id
}
