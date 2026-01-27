resource "cloudflare_record" "%[1]s" {
  zone_id = "%[2]s"
  name    = "%[3]s"
  proxied = false
  ttl     = 1
  type    = "TXT"
  value   = "v=spf1 -all"
}
