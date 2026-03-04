resource "cloudflare_zero_trust_device_default_profile" "%s" {
  account_id                     = "%s"
  register_interface_ip_with_dns = true
  sccm_vpn_boundary_support      = false

  allow_mode_switch = true
  auto_connect      = 0
  captive_portal    = 300
  tunnel_protocol   = "wireguard"

  service_mode_v2 = {
    mode = "proxy"
    port = 8080
  }
}
