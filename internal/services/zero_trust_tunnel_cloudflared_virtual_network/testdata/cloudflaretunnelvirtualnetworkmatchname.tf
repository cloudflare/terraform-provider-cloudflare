
resource "cloudflare_tunnel_virtual_network" "%[2]s" {
	account_id = "%[1]s"
	name       = "%[2]s"
	comment     = "test"
}
data "cloudflare_tunnel_virtual_network" "%[2]s" {
	account_id = cloudflare_tunnel_virtual_network.%[2]s.account_id
	name       = cloudflare_tunnel_virtual_network.%[2]s.name
	depends_on = ["cloudflare_tunnel_virtual_network.%[2]s"]
}
