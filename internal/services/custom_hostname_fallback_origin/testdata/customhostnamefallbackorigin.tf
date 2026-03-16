# DNS record must be created before the fallback origin can reference it.
# This ensures correct destruction order: fallback origin deleted first, then DNS record.
resource "cloudflare_dns_record" "%[2]s" {
  zone_id = "%[1]s"
  name    = "fallback-origin.%[2]s.%[4]s"
  content = "example.com"
  type    = "CNAME"
  proxied = true
  ttl     = 1
}

resource "cloudflare_custom_hostname_fallback_origin" "%[2]s" {
  zone_id    = "%[1]s"
  origin     = "fallback-origin.%[3]s.%[4]s"
  depends_on = [cloudflare_dns_record.%[2]s]
}
