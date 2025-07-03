
  resource "cloudflare_magic_wan_gre_tunnel" "%[1]s" {
	account_id = "%[4]s"
	name = "%[2]s"
	customer_gre_endpoint = "203.0.113.1"
	cloudflare_gre_endpoint = "%[5]s"
	interface_address = "10.213.0.9/31"
	description = "%[3]s"
  }