resource "cloudflare_ipsec_tunnel" "example" {
  account_id           = "c4a7362d577a6c3019a474fd6f485821"
  name                 = "IPsec_1"
  customer_endpoint    = "203.0.113.1"
  cloudflare_endpoint  = "203.0.113.1"
  interface_address    = "192.0.2.0/31"
  description          = "Tunnel for ISP X"
  health_check_enabled = true
  health_check_target  = "203.0.113.1"
  health_check_type    = "reply"
  psk                  = "asdf12341234"
}
