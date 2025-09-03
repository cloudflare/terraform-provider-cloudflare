
  resource "cloudflare_magic_wan_gre_tunnel" "%[1]s" {
	account_id = "%[4]s"
	name = "%[2]s"
	customer_gre_endpoint = "%[6]s"
	cloudflare_gre_endpoint = "%[5]s"
	interface_address = "%[7]s"
	description = "%[3]s"
  ttl = 65
  mtu = 1475
  health_check = {
    enabled = false
    }
  }