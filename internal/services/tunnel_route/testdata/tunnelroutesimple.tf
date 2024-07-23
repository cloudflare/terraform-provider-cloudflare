
resource "cloudflare_tunnel" "%[1]s" {
	account_id = "%[3]s"
	name       = "%[1]s"
	secret     = "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="
}

resource "cloudflare_tunnel_route" "%[1]s" {
    account_id = "%[3]s"
    tunnel_id = cloudflare_tunnel.%[1]s.id
    network = "%[4]s"
    comment = "%[2]s"
}