
  resource "cloudflare_magic_wan_ipsec_tunnel" "temp_ipsec_tunnel" {
	account_id = "%[2]s"
	name = "%[1]s"
	customer_endpoint = "203.0.113.1"
	cloudflare_endpoint = "%[4]s"
	interface_address = "%[5]s"
	description = ""
	health_check = {
		enabled = true
		rate = "low"
		type = "request"
		direction = "unidirectional"
	}
	psk = "abcde"
	replay_protection = true
  }
  
  resource "cloudflare_magic_wan_static_route" "%[1]s" {
	account_id = "%[2]s"
	prefix = "10.100.0.0/24"
	nexthop = "%[6]s"
	priority = "100"
	weight = %[3]d
	depends_on  = [cloudflare_magic_wan_ipsec_tunnel.temp_ipsec_tunnel]
  }
