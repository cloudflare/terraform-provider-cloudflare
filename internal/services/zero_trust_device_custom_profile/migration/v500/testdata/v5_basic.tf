resource "cloudflare_zero_trust_device_custom_profile" "%s" {
  account_id            = "%s"
  name                  = "Custom Profile Test"
  description           = "Test custom device profile"
  match                 = "identity.email == \"test@example.com\""
  precedence            = %d
  disable_auto_fallback = false
  captive_portal        = 180
  allow_mode_switch     = false
  switch_locked         = false
  allow_updates         = true
  auto_connect          = 0
  allowed_to_leave      = true
  support_url           = "https://support.example.com"
  service_mode_v2 = {
    mode = "proxy"
    port = 3128
  }
  tunnel_protocol       = "wireguard"
}
