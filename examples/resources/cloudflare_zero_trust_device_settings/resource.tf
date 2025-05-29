resource "cloudflare_zero_trust_device_settings" "example_zero_trust_device_settings" {
  account_id = "699d98642c564d2e855e9661899b7252"
  disable_for_time = 0
  gateway_proxy_enabled = true
  gateway_udp_proxy_enabled = true
  root_certificate_installation_enabled = true
  use_zt_virtual_ip = true
}
