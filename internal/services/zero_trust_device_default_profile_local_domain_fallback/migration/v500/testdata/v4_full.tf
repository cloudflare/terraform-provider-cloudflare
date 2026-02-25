resource "cloudflare_zero_trust_local_fallback_domain" "%[1]s" {
  account_id = "%[2]s"

  domains {
    suffix      = "example.com"
    description = "Primary domain"
    dns_server  = ["1.1.1.1", "1.0.0.1"]
  }

  domains {
    suffix      = "internal.example.com"
    description = "Internal services"
  }
}
