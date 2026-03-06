resource "cloudflare_zero_trust_device_profiles" "%s" {
  account_id  = "%s"
  name        = "Test Profile"
  description = "Test device profile for migration"
  default     = true

  allow_mode_switch    = false
  auto_connect         = 0
  captive_portal       = 180
  disable_auto_fallback = false
  switch_locked        = false
  allow_updates        = true
  allowed_to_leave     = true
  support_url          = "https://support.example.com"
  tunnel_protocol      = "wireguard"
  service_mode_v2_mode = "proxy"
  service_mode_v2_port = 8080
}
