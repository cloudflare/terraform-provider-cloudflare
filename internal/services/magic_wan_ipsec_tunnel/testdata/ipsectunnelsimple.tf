
  resource "cloudflare_magic_wan_ipsec_tunnel" "%[1]s" {
	account_id = "%[3]s"
	name = "%[2]s"
	customer_endpoint = "203.0.113.1"
	cloudflare_endpoint = "%[5]s"
	interface_address = "%[6]s"
	description = "%[2]s"
	health_check = {
		enabled = true
		rate = "low"
		type = "request"
		direction = "unidirectional"
	}
	psk = "%[4]s"
	replay_protection = true
	automatic_return_routing = true
	bgp = {
		customer_asn = 65001
	}
  }