resource "cloudflare_dns_record" "%[3]s" {
  zone_id = "%[1]s"
  name = "%[2]s.%[3]s.%[4]s"
  type    = "CNAME"
  content = "dev.%[2]s.%[3]s.%[4]s"
  ttl     = 1
  proxied = false
  comment = "Test record"
  settings = {}
}