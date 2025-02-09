
  resource "cloudflare_magic_wan_gre_tunnel" "%[1]s" {
	account_id = "%[4]s"
	name = "%[2]s"
	customer_gre_endpoint = "203.0.113.1"
	cloudflare_gre_endpoint = "162.159.64.41"
	interface_address = "10.212.0.9/31"
	description = "%[3]s"
    ttl = 64
    mtu = 1476
    health_check_enabled = true
    health_check_target = "203.0.113.1"
    health_check_type = "request"
  }