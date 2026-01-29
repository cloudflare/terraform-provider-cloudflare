resource "cloudflare_record" "%[1]s" {
  zone_id = "%[2]s"
  name    = "%[3]s"
  type    = "PTR"
  value   = "example.com"
  ttl     = 3600
}
