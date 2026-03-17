resource "cloudflare_zero_trust_local_fallback_domain" "%[1]s" {
  account_id = "%[2]s"
  policy_id  = null

  domains {
    suffix      = "test.example.com"
    description = "Test domain with null policy_id"
    dns_server  = ["10.2.0.1"]
  }
}
