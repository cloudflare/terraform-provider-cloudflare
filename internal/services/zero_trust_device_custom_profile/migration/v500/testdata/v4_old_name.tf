resource "cloudflare_device_settings_policy" "%s" {
  account_id  = "%s"
  name        = "Old Name Custom Profile"
  description = "Test custom profile with old resource name"
  match       = "identity.email == \"legacy@example.com\""
  precedence  = %d

  allow_mode_switch    = true
  auto_connect         = 0
  captive_portal       = 300
  service_mode_v2_mode = "proxy"
  service_mode_v2_port = 8080
}
