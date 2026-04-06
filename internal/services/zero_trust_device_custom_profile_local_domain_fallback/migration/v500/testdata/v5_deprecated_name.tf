resource "cloudflare_zero_trust_device_custom_profile" "%[1]s_profile" {
  account_id  = "%[2]s"
  name        = "Deprecated Name Profile"
  description = "Using deprecated resource name"
  match       = "identity.email == \"legacy@example.com\""
  precedence  = %[3]d
}

moved {
  from = cloudflare_device_settings_policy.%[1]s_profile
  to   = cloudflare_zero_trust_device_custom_profile.%[1]s_profile
}

resource "cloudflare_zero_trust_device_custom_profile_local_domain_fallback" "%[1]s" {
  account_id = "%[2]s"
  policy_id  = cloudflare_zero_trust_device_custom_profile.%[1]s_profile.id

  domains = [
    {
      suffix      = "deprecated.example.com"
      description = "Using deprecated fallback domain name"
    }
  ]
}

moved {
  from = cloudflare_fallback_domain.%[1]s
  to   = cloudflare_zero_trust_device_custom_profile_local_domain_fallback.%[1]s
}
