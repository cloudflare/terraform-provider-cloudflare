
  resource "cloudflare_magic_wan_ipsec_tunnel" "%[1]s" {
	account_id = "%[3]s"
	name = "%[2]s"
	customer_endpoint = "203.0.113.1"
	cloudflare_endpoint = "162.159.64.41"
	interface_address = "10.212.0.9/31"
	description = "%[2]s"
	health_check_enabled = true
	health_check_target = "203.0.113.1"
	health_check_type = "request"
	health_check_direction = "unidirectional"
	health_check_rate = "mid"
	psk = "%[4]s"
	allow_null_cipher = false
	replay_protection = true
  }