
  resource "cloudflare_magic_wan_ipsec_tunnel" "%[1]s" {
	account_id = "%[3]s"
	name = "%[2]s"
	customer_endpoint = "203.0.113.1"
	cloudflare_endpoint = "%[5]s"
	interface_address = "10.212.0.9/31"
	description = "%[2]s"
	health_check = {
		enabled = true
		rate = "low"
		type = "request"
		direction = "unidirectional"
	}
	psk = "%[4]s"
	replay_protection = true
  }