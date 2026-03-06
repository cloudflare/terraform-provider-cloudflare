resource "cloudflare_zero_trust_device_profiles" "%s" {
  account_id  = "%s"
  name        = "Maximal Custom Profile"
  description = "Test custom profile with all optional fields"
  match       = "identity.email == \"maximal@example.com\""
  precedence  = %d

  allow_mode_switch     = false
  allow_updates         = true
  allowed_to_leave      = false
  auto_connect          = 30
  captive_portal        = 600
  disable_auto_fallback = true
  switch_locked         = true
  support_url           = "https://support.custom.cf-tf-test.com"

  service_mode_v2_mode = "proxy"
  service_mode_v2_port = 443
}
