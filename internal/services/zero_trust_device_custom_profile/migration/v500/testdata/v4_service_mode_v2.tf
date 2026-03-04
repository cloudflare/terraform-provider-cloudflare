resource "cloudflare_zero_trust_device_profiles" "%s" {
  account_id  = "%s"
  name        = "Custom Profile with Service Mode"
  description = "Test custom profile with service_mode_v2 fields"
  match       = "identity.email == \"admin@example.com\""
  precedence  = %d

  allow_mode_switch    = false
  auto_connect         = 15
  captive_portal       = 300
  service_mode_v2_mode = "warp"
}
