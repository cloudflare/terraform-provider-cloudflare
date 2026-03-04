resource "cloudflare_zero_trust_device_default_profile_local_domain_fallback" "%[1]s" {
  account_id = "%[2]s"

  domains = [
    {
      suffix      = "corp.example.com"
      description = "Corporate network"
      dns_server  = ["10.0.0.1", "10.0.0.2"]
    },
    {
      suffix      = "internal.example.com"
      description = "Internal services"
      dns_server  = ["10.1.0.1"]
    },
    {
      suffix = "local.example.com"
    }
  ]
}

moved {
  from = cloudflare_zero_trust_local_fallback_domain.%[1]s
  to   = cloudflare_zero_trust_device_default_profile_local_domain_fallback.%[1]s
}
