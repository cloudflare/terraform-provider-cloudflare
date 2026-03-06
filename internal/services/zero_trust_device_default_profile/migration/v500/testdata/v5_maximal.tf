resource "cloudflare_zero_trust_device_default_profile" "%s" {
  account_id                     = "%s"
  register_interface_ip_with_dns = true
  sccm_vpn_boundary_support      = false

  allow_mode_switch     = false
  allow_updates         = true
  allowed_to_leave      = true
  auto_connect          = 30
  captive_portal        = 600
  disable_auto_fallback = true
  switch_locked         = true
  support_url           = "https://support.cf-tf-test.com"
  tunnel_protocol       = "wireguard"

  service_mode_v2 = {
    mode = "proxy"
    port = 443
  }
}
