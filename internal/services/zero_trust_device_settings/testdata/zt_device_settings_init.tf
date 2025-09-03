resource "cloudflare_zero_trust_device_settings" "%[1]s" {
  account_id = "%[2]s"
  disable_for_time = 20
  gateway_proxy_enabled = true
  gateway_udp_proxy_enabled = true
  root_certificate_installation_enabled = true
  use_zt_virtual_ip = true
}
