resource "cloudflare_zero_trust_device_custom_profile" "%[1]s_profile" {
  account_id  = "%[2]s"
  name        = "Test Custom Profile Full"
  description = "Custom profile with all fields"
  match       = "identity.email == \"admin@example.com\""
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
      suffix      = "example.com"
      description = "Primary domain"
      dns_server  = ["1.1.1.1", "1.0.0.1"]
    },
    {
      suffix      = "internal.example.com"
      description = "Internal services"
    }
  ]
}

moved {
  from = cloudflare_zero_trust_local_fallback_domain.%[1]s
  to   = cloudflare_zero_trust_device_custom_profile_local_domain_fallback.%[1]s
}
