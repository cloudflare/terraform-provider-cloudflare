resource "cloudflare_dns_record" "%[1]s" {
  zone_id = "%[2]s"
  name    = "%[3]s"
  type    = "PTR"
  content = "example.com"
  ttl     = 3600
}
