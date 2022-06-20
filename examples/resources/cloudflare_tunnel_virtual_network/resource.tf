resource "cloudflare_tunnel_virtual_network" "example" {
  account_id = "c4a7362d577a6c3019a474fd6f485821"
  name = "vnet-for-documentation"
  comment = "New tunnel virtual network for documentation"
}