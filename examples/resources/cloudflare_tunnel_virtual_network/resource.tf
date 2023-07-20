resource "cloudflare_tunnel_virtual_network" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name = "vnet-for-documentation"
  comment = "New tunnel virtual network for documentation"
}