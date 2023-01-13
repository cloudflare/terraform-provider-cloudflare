resource "cloudflare_gre_tunnel" "example" {
  account_id              = "f037e56e89293a057740de681ac9abbe"
  name                    = "GRE_1"
  customer_gre_endpoint   = "203.0.113.1"
  cloudflare_gre_endpoint = "203.0.113.1"
  interface_address       = "192.0.2.0/31"
  description             = "Tunnel for ISP X"
  ttl                     = 64
  mtu                     = 1476
  health_check_enabled    = true
  health_check_target     = "203.0.113.1"
  health_check_type       = "reply"
}
