resource "cloudflare_device_settings_policy" "%[1]s_profile" {
  account_id  = "%[2]s"
  name        = "Deprecated Name Profile"
  description = "Using deprecated resource name"
  match       = "identity.email == \"legacy@example.com\""
  precedence  = %[3]d
}

resource "cloudflare_fallback_domain" "%[1]s" {
  account_id = "%[2]s"
  policy_id  = cloudflare_device_settings_policy.%[1]s_profile.id

  domains {
    suffix      = "deprecated.example.com"
    description = "Using deprecated fallback domain name"
  }
}
