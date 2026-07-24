resource "cloudflare_dns_record" "%[2]s" {
  zone_id = "%[1]s"
  name    = "_https.%[2]s.%[3]s"
  type    = "URI"
  ttl     = 300

  data = {
    priority = %[4]s
    weight   = 10
    target   = "https://example.com/"
  }
}
