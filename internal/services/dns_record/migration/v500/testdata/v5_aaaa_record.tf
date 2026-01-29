resource "cloudflare_dns_record" "%[1]s" {
  zone_id = "%[2]s"
  name    = "%[3]s"
  type    = "AAAA"
  content = "2001:db8::1"
  ttl     = 3600
}
