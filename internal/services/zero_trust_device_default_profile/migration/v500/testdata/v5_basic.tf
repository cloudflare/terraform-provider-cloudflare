resource "cloudflare_zero_trust_device_default_profile" "%s" {
  account_id                     = "%s"
  tunnel_protocol                = "wireguard"
  register_interface_ip_with_dns = true
  sccm_vpn_boundary_support      = false

  allow_mode_switch     = false
  auto_connect          = 0
  captive_portal        = 180
  disable_auto_fallback = false
  switch_locked         = false
  allow_updates         = true
  allowed_to_leave      = true
  support_url           = "https://support.example.com"

  service_mode_v2 = {
    mode = "proxy"
    port = 8080
  }
}
