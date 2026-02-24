resource "cloudflare_tunnel" "%[1]s" {
  account_id = "%[2]s"
  name       = "tf-acc-test-tunnel-or-%[1]s"
  secret     = "%[3]s"
  config_src = "cloudflare"
}

resource "cloudflare_zero_trust_tunnel_cloudflared_config" "%[1]s" {
  account_id = "%[2]s"
  tunnel_id  = cloudflare_tunnel.%[1]s.id

  config {
    origin_request {
      no_happy_eyeballs      = false
      keep_alive_connections = 1024
      http_host_header       = "example.internal"
      http2_origin           = true
    }

    ingress_rule {
      hostname = "test.example.com"
      service  = "http://localhost:8080"
    }

    ingress_rule {
      service = "http_status:404"
    }
  }
}
