
resource "cloudflare_tunnel_virtual_network" "%[1]s" {
	account_id         = "%[3]s"
	name               = "%[4]s"
	comment            = "%[2]s"
	is_default_network = "%[5]t"
}