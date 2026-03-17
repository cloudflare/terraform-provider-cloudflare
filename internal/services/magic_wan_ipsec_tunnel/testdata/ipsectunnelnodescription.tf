
  resource "cloudflare_magic_wan_ipsec_tunnel" "%[1]s" {
	account_id = "%[2]s"
	name = "%[1]s"
	customer_endpoint = "203.0.113.1"
	cloudflare_endpoint = "%[3]s"
	interface_address = "%[4]s"
	health_check = {
		enabled = true
		rate = "low"
		type = "request"
		direction = "unidirectional"
	}
	psk = "%[5]s"
	replay_protection = true
	automatic_return_routing = true
	bgp = {
		customer_asn = 65001
	}
  }
