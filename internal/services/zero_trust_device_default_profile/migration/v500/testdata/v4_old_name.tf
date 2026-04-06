resource "cloudflare_device_settings_policy" "%s" {
  account_id  = "%s"
  name        = "Old Name Test Profile"
  description = "Test device profile with old resource name"
  default     = true

  allow_mode_switch    = true
  auto_connect         = 0
  captive_portal       = 300
  tunnel_protocol      = "wireguard"
  service_mode_v2_mode = "proxy"
  service_mode_v2_port = 8080
}
