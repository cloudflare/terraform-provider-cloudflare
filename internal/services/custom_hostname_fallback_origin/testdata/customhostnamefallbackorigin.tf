
resource "cloudflare_custom_hostname_fallback_origin" "%[2]s" {
  zone_id = "%[1]s"
  origin = "fallback-origin.%[3]s.%[4]s"
}

resource "cloudflare_dns_record" "%[2]s" {
  zone_id = "%[1]s"
  name    = "fallback-origin.%[2]s.%[4]s"
  content   = "example.com"
  type    = "CNAME"
  proxied = true
  ttl     = 1
  depends_on = [cloudflare_custom_hostname_fallback_origin.%[2]s]
}
