resource "cloudflare_zero_trust_tunnel_cloudflared_virtual_network" "%[1]s" {
  account_id         = "%[2]s"
  name               = "tf-acc-test-vnet-complete-%[1]s"
  comment            = "Migration test complete"
  is_default_network = false
}
