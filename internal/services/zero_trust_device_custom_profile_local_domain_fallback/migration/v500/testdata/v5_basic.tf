resource "cloudflare_zero_trust_device_custom_profile" "%[1]s_profile" {
  account_id  = "%[2]s"
  name        = "Test Custom Profile Basic"
  description = "Custom profile for migration testing"
  match       = "identity.email == \"test@example.com\""
  precedence  = %[3]d
}

moved {
  from = cloudflare_zero_trust_device_profiles.%[1]s_profile
  to   = cloudflare_zero_trust_device_custom_profile.%[1]s_profile
}

resource "cloudflare_zero_trust_device_custom_profile_local_domain_fallback" "%[1]s" {
  account_id = "%[2]s"
  policy_id  = cloudflare_zero_trust_device_custom_profile.%[1]s_profile.id

  domains = [
    {
      suffix = "example.com"
    }
  ]
}

moved {
  from = cloudflare_zero_trust_local_fallback_domain.%[1]s
  to   = cloudflare_zero_trust_device_custom_profile_local_domain_fallback.%[1]s
}
