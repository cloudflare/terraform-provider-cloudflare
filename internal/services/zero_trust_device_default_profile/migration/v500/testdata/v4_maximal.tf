resource "cloudflare_zero_trust_device_profiles" "%s" {
  account_id  = "%s"
  name        = "Maximal Test Profile"
  description = "Test device profile with all optional fields"
  default     = true

  allow_mode_switch     = false
  allow_updates         = true
  allowed_to_leave      = true
  auto_connect          = 30
  captive_portal        = 600
  disable_auto_fallback = true
  switch_locked         = true
  support_url           = "https://support.cf-tf-test.com"
  tunnel_protocol       = "wireguard"

  service_mode_v2_mode = "proxy"
  service_mode_v2_port = 443
}
