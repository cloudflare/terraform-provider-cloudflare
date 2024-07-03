
resource "cloudflare_fallback_domain" "%[1]s" {
  account_id = "%[2]s"
  domains =[ {
    description = "%[3]s"
    suffix      = "%[4]s"
    dns_server  = ["%[5]s"]
  }]
}
