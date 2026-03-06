resource "cloudflare_zero_trust_device_default_profile_local_domain_fallback" "%[1]s" {
  account_id = "%[2]s"

  domains = [
    {
      suffix      = "internal.example.com"
      description = "Internal domain"
      dns_server  = ["10.1.0.1"]
    }
  ]
}

moved {
  from = cloudflare_fallback_domain.%[1]s
  to   = cloudflare_zero_trust_device_default_profile_local_domain_fallback.%[1]s
}
