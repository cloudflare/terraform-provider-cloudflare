resource "cloudflare_zero_trust_device_default_profile_local_domain_fallback" "%[1]s" {
  account_id = "%[2]s"

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
  to   = cloudflare_zero_trust_device_default_profile_local_domain_fallback.%[1]s
}
