
resource "cloudflare_zero_trust_tunnel_cloudflared" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  secret     = "%[3]s"
}

resource "cloudflare_zero_trust_tunnel_cloudflared_config" "%[1]s" {
  account_id = "%[2]s"
  tunnel_id  = cloudflare_zero_trust_tunnel_cloudflared.%[1]s.id

  config = {
    warp_routing = {
      enabled = true
    }
    origin_request = {
      connect_timeout          = "1m0s"
      tls_timeout              = "1m0s"
      tcp_keep_alive           = "1m0s"
      no_happy_eyeballs        = false
      keep_alive_connections   = 1024
      keep_alive_timeout       = "1m0s"
      http_host_header         = "baz"
      origin_server_name       = "foobar"
      ca_pool                  = "/path/to/unsigned/ca/pool"
      no_tls_verify            = false
      disable_chunked_encoding = false
      bastion_mode             = false
      proxy_address            = "10.0.0.1"
      proxy_port               = "8123"
      proxy_type               = "socks"
      http2_origin             = true
      ip_rules = [{
        prefix = "/web"
        ports  = [80, 443]
        allow  = false
      }]
    }
    ingress_rule = [{
      hostname = "foo"
      path     = "/bar"
      service  = "http://10.0.0.2:8080"
      origin_request = {
        connect_timeout = "15s"
      }
      },
      {
        hostname = "bar"
        path     = "/foo"
        service  = "http://10.0.0.3:8081"
        origin_request = {
          access = {
            required  = true
            team_name = "terraform"
            aud_tag   = ["AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"]
          }
        }
      },
      {
        service = "https://10.0.0.4:8082"
      }
    ]
  }
}
