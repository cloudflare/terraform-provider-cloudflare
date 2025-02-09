
resource "cloudflare_zero_trust_tunnel_cloudflared" "%[2]s" {
	account_id = "%[1]s"
	name       = "%[2]s"
	tunnel_secret     = "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="
}

data "cloudflare_zero_trust_tunnel_cloudflared" "%[2]s" {
	account_id = cloudflare_zero_trust_tunnel_cloudflared.%[2]s.account_id
	name       = cloudflare_zero_trust_tunnel_cloudflared.%[2]s.name
	depends_on = [cloudflare_zero_trust_tunnel_cloudflared.%[2]s]
}
