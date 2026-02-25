resource "cloudflare_fallback_domain" "%[1]s" {
  account_id = "%[2]s"

  domains {
    suffix      = "internal.example.com"
    description = "Internal domain"
    dns_server  = ["10.1.0.1"]
  }
}
