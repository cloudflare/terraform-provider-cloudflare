resource "cloudflare_zero_trust_tunnel_cloudflared" "%[2]s" {
	account_id    = "%[1]s"
	name          = "%[2]s"
	tunnel_secret = "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="
}

resource "cloudflare_zero_trust_tunnel_cloudflared" "%[3]s" {
	account_id    = "%[1]s"
	name          = "%[3]s"
	tunnel_secret = "UGBAECAwQFwgBAgIDBAUMEAQIDBABQYHCBgcIAQGBwg="
}

resource "cloudflare_zero_trust_tunnel_cloudflared_route" "%[4]s" {
    account_id = "%[1]s"
    tunnel_id = cloudflare_zero_trust_tunnel_cloudflared.%[5]s.id
    network = "1.1.1.1/32"
}
