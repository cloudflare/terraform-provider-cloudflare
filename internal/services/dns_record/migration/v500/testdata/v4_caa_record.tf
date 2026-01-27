resource "cloudflare_record" "%[1]s" {
  zone_id = "%[2]s"
  name    = "%[3]s"
  proxied = false
  ttl     = 1
  type    = "CAA"

  data {
    flags = 0
    tag   = "issue"
    value = "letsencrypt.org"
  }
}
