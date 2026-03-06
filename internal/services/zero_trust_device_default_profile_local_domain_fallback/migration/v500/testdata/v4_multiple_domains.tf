resource "cloudflare_zero_trust_local_fallback_domain" "%[1]s" {
  account_id = "%[2]s"

  domains {
    suffix      = "corp.example.com"
    description = "Corporate network"
    dns_server  = ["10.0.0.1", "10.0.0.2"]
  }

  domains {
    suffix      = "internal.example.com"
    description = "Internal services"
    dns_server  = ["10.1.0.1"]
  }

  domains {
    suffix = "local.example.com"
  }
}
