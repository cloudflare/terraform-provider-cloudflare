resource "cloudflare_zero_trust_device_custom_profile" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[1]s"
  match       = "identity.email == \"test@example.com\""
  precedence  = 100
  enabled     = true
  description = "Test custom device profile"
}