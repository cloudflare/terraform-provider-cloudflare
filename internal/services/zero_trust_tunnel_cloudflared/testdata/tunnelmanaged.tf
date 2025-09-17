
	resource "cloudflare_zero_trust_tunnel_cloudflared" "%[2]s" {
		account_id = "%[1]s"
		name       = "%[2]s"
		tunnel_secret     = "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="
		config_src = "cloudflare"
	}
