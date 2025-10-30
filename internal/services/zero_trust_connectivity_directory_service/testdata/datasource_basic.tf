resource "cloudflare_zero_trust_tunnel_cloudflared" "%[1]s" {
	account_id    = "%[2]s"
	name          = "%[1]s"
	config_src    = "cloudflare"
	tunnel_secret = "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="
}

resource "cloudflare_zero_trust_connectivity_directory_service" "%[1]s" {
	account_id = "%[2]s"
	name       = "%[3]s"
	type       = "http"
	host = {
		hostname = "%[4]s"
		resolver_network = {
			tunnel_id    = cloudflare_zero_trust_tunnel_cloudflared.%[1]s.id
			resolver_ips = ["%[5]s", "%[6]s"]
		}
	}
	http_port  = %[7]d
	https_port = %[8]d
}

data "cloudflare_zero_trust_connectivity_directory_service" "%[1]s" {
	account_id = "%[2]s"
	service_id = cloudflare_zero_trust_connectivity_directory_service.%[1]s.service_id
}