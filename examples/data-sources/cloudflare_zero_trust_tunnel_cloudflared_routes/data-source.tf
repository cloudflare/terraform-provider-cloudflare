data "cloudflare_zero_trust_tunnel_cloudflared_routes" "example_zero_trust_tunnel_cloudflared_routes" {
  account_id = "699d98642c564d2e855e9661899b7252"
  comment = "Example comment for this route."
  existed_at = "2019-10-12T07:20:50.52Z"
  is_deleted = true
  network_subset = "172.16.0.0/16"
  network_superset = "172.16.0.0/16"
  route_id = "f70ff985-a4ef-4643-bbbc-4a0ed4fc8415"
  tun_types = "cfd_tunnel,warp_connector"
  tunnel_id = "f70ff985-a4ef-4643-bbbc-4a0ed4fc8415"
  virtual_network_id = "f70ff985-a4ef-4643-bbbc-4a0ed4fc8415"
}
