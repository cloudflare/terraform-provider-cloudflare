# Tunnel route
resource "cloudflare_tunnel_route" "example" {
  account_id         = "c4a7362d577a6c3019a474fd6f485821"
  tunnel_id          = "f70ff985-a4ef-4643-bbbc-4a0ed4fc8415"
  network            = "192.0.2.24/32"
  comment            = "New tunnel route for documentation"
  virtual_network_id = "bdc39a3c-3104-4c23-8ac0-9f455dda691a"
}

# Tunnel with tunnel route
resource "cloudflare_argo_tunnel" "tunnel" {
  account_id = "c4a7362d577a6c3019a474fd6f485821"
  name       = "my_tunnel"
  secret     = "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="
}

resource "cloudflare_tunnel_route" "example" {
  account_id         = "c4a7362d577a6c3019a474fd6f485821"
  tunnel_id          = cloudflare_argo_tunnel.tunnel.id
  network            = "192.0.2.24/32"
  comment            = "New tunnel route for documentation"
  virtual_network_id = "bdc39a3c-3104-4c23-8ac0-9f455dda691a"
}
