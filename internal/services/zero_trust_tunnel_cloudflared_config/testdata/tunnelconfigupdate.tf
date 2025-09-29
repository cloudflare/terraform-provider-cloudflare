resource "cloudflare_zero_trust_tunnel_cloudflared" "%[1]s" {
  account_id    = "%[2]s"
  name          = "%[1]s"
  tunnel_secret = "%[3]s"
}

resource "cloudflare_zero_trust_tunnel_cloudflared_config" "%[1]s" {
  account_id         = "%[2]s"
  tunnel_id          = cloudflare_zero_trust_tunnel_cloudflared.%[1]s.id

  config = {
    origin_request = {
      connect_timeout          = 61
      tls_timeout              = 61
      tcp_keep_alive           = 61
      no_happy_eyeballs        = true
      keep_alive_connections   = 1028
      keep_alive_timeout       = 61
      http_host_header         = "bez"
      origin_server_name       = "fuuber"
      ca_pool                  = "/path/to/unsigned/ca/pool"
      no_tls_verify            = false
      disable_chunked_encoding = false
      proxy_type               = "socks"
      http2_origin             = true
    }

    ingress = [
      {
        hostname = "fuu"
        path     = "/ber"
        service  = "http://10.0.0.3:8080"
        origin_request = {
          connect_timeout = 20
        }
      },
      {
        hostname = "ber"
        path     = "/fuu"
        service  = "http://10.0.0.5:8081"
        origin_request = {
          access = {
            required  = false
            team_name = "terraform2"
            aud_tag   = ["BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB"]
          }
        }
      },
      {
        service = "https://10.0.0.5:8082"
      }
    ]
  }
}
