resource "cloudflare_tunnel_virtual_network" "%[1]s" {
  account_id = "%[2]s"
  name       = "tf-acc-test-vnet-%[1]s"
}
