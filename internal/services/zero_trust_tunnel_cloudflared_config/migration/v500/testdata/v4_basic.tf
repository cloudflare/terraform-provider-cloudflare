resource "cloudflare_tunnel" "%[1]s" {
  account_id = "%[2]s"
  name       = "tf-acc-test-tunnel-%[1]s"
  secret     = "%[3]s"
  config_src = "cloudflare"
}

resource "cloudflare_zero_trust_tunnel_cloudflared_config" "%[1]s" {
  account_id = "%[2]s"
  tunnel_id  = cloudflare_tunnel.%[1]s.id

  config {
    ingress_rule {
      service = "http_status:404"
    }
  }
}
