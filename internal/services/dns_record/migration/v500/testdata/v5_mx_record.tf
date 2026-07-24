resource "cloudflare_dns_record" "%[1]s" {
  zone_id  = "%[2]s"
  name     = "%[3]s"
  type     = "MX"
  data = {
    priority = 10
    target   = "mail.example.com"
  }
  ttl      = 1
}
