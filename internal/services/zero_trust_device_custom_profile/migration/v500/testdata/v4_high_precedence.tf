resource "cloudflare_zero_trust_device_profiles" "%s" {
  account_id  = "%s"
  name        = "High Priority Profile"
  description = "Custom profile with high precedence"
  match       = "identity.email == \"vip@example.com\""
  precedence  = %d

  allow_mode_switch    = true
  auto_connect         = 5
  captive_portal       = 120
  service_mode_v2_mode = "proxy"
  service_mode_v2_port = 8080
}
