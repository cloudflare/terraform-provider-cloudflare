resource "cloudflare_dns_record" "%[1]s" {
  zone_id = "%[2]s"
  name    = "selector2._domainkey.%[1]s.%[3]s"
  content = "selector2-test._domainkey.MixedCase.onmicrosoft.com"
  type    = "CNAME"
  proxied = false
  ttl     = 60
}
