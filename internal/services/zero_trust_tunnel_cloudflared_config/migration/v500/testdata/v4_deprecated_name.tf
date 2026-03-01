resource "cloudflare_tunnel" "%[1]s" {
  account_id = "%[2]s"
  name       = "tf-acc-test-tunnel-dep-%[1]s"
  secret     = "%[3]s"
  config_src = "cloudflare"
}

resource "cloudflare_tunnel_config" "%[1]s" {
  account_id = "%[2]s"
  tunnel_id  = cloudflare_tunnel.%[1]s.id

  config {
    ingress_rule {
      service = "http_status:404"
    }
  }
}
