
resource "cloudflare_zero_trust_tunnel_cloudflared" "%[1]s" {
	account_id    = "%[3]s"
	name          = "%[1]s"
	tunnel_secret = "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="
}

resource "cloudflare_zero_trust_tunnel_cloudflared_route" "%[1]s" {
    account_id = "%[3]s"
    tunnel_id = cloudflare_zero_trust_tunnel_cloudflared.%[1]s.id
    network = "%[4]s"
    comment = "%[2]s"
}
