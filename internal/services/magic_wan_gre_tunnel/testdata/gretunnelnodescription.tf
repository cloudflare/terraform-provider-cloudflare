
  resource "cloudflare_magic_wan_gre_tunnel" "%[1]s" {
	account_id = "%[3]s"
	name = "%[2]s"
	customer_gre_endpoint = "%[4]s"
	cloudflare_gre_endpoint = "%[5]s"
	interface_address = "%[6]s"
	automatic_return_routing = true
	bgp = {
		customer_asn = 65002
	}
  }
