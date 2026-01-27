resource "cloudflare_record" "%[1]s" {
  zone_id = "%[2]s"
  name    = "%[3]s"
  type    = "SRV"
  ttl     = 3600

  data {
    priority = 10
    weight   = 60
    port     = 5060
    target   = "sipserver.example.com"
  }
}
