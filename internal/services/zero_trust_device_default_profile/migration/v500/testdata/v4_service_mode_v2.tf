resource "cloudflare_zero_trust_device_profiles" "%s" {
  account_id  = "%s"
  name        = "Test Profile with Service Mode"
  description = "Test device profile with service_mode_v2 fields"
  default     = true

  allow_mode_switch    = true
  auto_connect         = 15
  captive_portal       = 300
  tunnel_protocol      = "wireguard"
  service_mode_v2_mode = "proxy"
  service_mode_v2_port = 8080
}
