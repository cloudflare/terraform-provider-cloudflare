
  resource "cloudflare_magic_wan_gre_tunnel" "%[1]s" {
	account_id = "%[4]s"
	name = "%[2]s"
	customer_gre_endpoint = "203.0.113.2"
	cloudflare_gre_endpoint = "%[5]s"
	interface_address = "10.213.0.11/31"
	description = "%[3]s"
  ttl = 65
  mtu = 1475
  health_check = {
    enabled = false
    }
  }