resource "cloudflare_dns_record" "%[1]s" {
  zone_id = "%[2]s"
  name    = "%[3]s"
  type    = "A"
  content = "192.0.2.3"
  ttl     = 3600
  tags    = ["env:test", "managed:terraform"]
}
