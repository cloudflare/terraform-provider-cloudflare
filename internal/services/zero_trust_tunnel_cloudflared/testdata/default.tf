
resource "cloudflare_tunnel" "%[2]s" {
	account_id = "%[1]s"
	name       = "%[2]s"
	secret     = "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="
}

data "cloudflare_tunnel" "%[2]s" {
	account_id = "%[1]s"
	name       = "%[2]s"
	is_deleted = false
	depends_on = [cloudflare_tunnel.%[2]s]
}
