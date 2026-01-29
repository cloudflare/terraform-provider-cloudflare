resource "cloudflare_record" "%[1]s" {
  zone_id = "%[2]s"
  name    = "%[3]s"
  type    = "AAAA"
  value   = "2001:db8::1"
  ttl     = 3600
}
