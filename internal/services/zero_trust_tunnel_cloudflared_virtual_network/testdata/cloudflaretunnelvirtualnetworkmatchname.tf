
resource "cloudflare_zero_trust_tunnel_cloudflared_virtual_network" "%[2]s" {
	account_id = "%[1]s"
	name       = "%[2]s"
	comment     = "test"
}
data "cloudflare_zero_trust_tunnel_cloudflared_virtual_network" "%[2]s" {
	account_id = cloudflare_zero_trust_tunnel_cloudflared_virtual_network.%[2]s.account_id
	name       = cloudflare_zero_trust_tunnel_cloudflared_virtual_network.%[2]s.name
	depends_on = ["cloudflare_zero_trust_tunnel_cloudflared_virtual_network.%[2]s"]
}
