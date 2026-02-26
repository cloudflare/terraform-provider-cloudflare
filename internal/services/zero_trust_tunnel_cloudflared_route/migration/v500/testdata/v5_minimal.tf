resource "cloudflare_zero_trust_tunnel_cloudflared_route" "%[1]s" {
  account_id = "%[2]s"
  tunnel_id  = "%[3]s"
  network    = "%[4]s"
}
