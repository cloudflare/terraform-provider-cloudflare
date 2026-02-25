resource "cloudflare_zero_trust_local_fallback_domain" "%[1]s" {
  account_id = "%[2]s"

  domains {
    suffix = "example.com"
  }
}
