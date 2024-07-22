
resource "cloudflare_tunnel_virtual_network" "%[2]s" {
	account_id = "%[1]s"
	name       = "my-tf-vnet-for-pool-%[2]s"
	comment     = "test"
}

resource "cloudflare_tunnel" "%[3]s" {
	account_id = "%[1]s"
	name       = "my-tf-tunnel-for-pool-%[3]s"
	secret     = "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="
}

resource "cloudflare_tunnel_route" "%[4]s" {
    account_id = "%[1]s"
    network = "192.0.2.1/32"
    virtual_network_id = cloudflare_tunnel_virtual_network.%[2]s.id
    tunnel_id = cloudflare_tunnel.%[3]s.id
    comment = "test"
}

resource "cloudflare_load_balancer_pool" "%[5]s" {
  account_id = "%[1]s"
  name = "my-tf-pool-with-vnet-%[5]s"
  latitude = 12.3
  longitude = 55
  origins =[ {
    name = "example-1"
    address = "192.0.2.1"
    virtual_network_id = cloudflare_tunnel_route.%[4]s.virtual_network_id
    enabled = true
  }
}]