
resource "cloudflare_dns_record" "%[1]s" {
  zone_id = "%[2]s"
  name = "%[3]s"
  data = {
  flags = "0"
    tag   = "issue"
    value = "letsencrypt.org"
  }
  type = "CAA"
  ttl = %[4]d
}
