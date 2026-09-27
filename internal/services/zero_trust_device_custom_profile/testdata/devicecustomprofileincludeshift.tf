resource "cloudflare_zero_trust_device_custom_profile" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  match      = "identity.email == \"nobody@example.com\""
  precedence = %[3]d
  enabled    = false

  include = %[4]s
}
