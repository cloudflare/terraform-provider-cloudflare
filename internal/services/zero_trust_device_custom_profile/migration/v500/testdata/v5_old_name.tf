resource "cloudflare_zero_trust_device_custom_profile" "%s" {
  account_id  = "%s"
  name        = "Old Name Custom Profile"
  description = "Test custom profile with old resource name"
  match       = "identity.email == \"legacy@example.com\""
  precedence  = %d

  allow_mode_switch = true
  auto_connect      = 0
  captive_portal    = 300

  service_mode_v2 = {
    mode = "proxy"
    port = 8080
  }
}
