
resource "cloudflare_zero_trust_tunnel_cloudflared_virtual_network" "%[2]s" {
	account_id = "%[1]s"
	name       = "%[2]s"
	comment     = "test"
  is_default_network = "false"
}

data "cloudflare_zero_trust_tunnel_cloudflared_virtual_network" "%[2]s" {
  virtual_network_id = cloudflare_zero_trust_tunnel_cloudflared_virtual_network.%[2]s.id
	account_id = cloudflare_zero_trust_tunnel_cloudflared_virtual_network.%[2]s.account_id
	depends_on = ["cloudflare_zero_trust_tunnel_cloudflared_virtual_network.%[2]s"]
}
