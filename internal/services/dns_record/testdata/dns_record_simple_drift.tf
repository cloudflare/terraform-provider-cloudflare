resource "cloudflare_dns_record" "%[1]s" {
  zone_id = "%[2]s"
  name    = "simple-drift.%[1]s.%[3]s"
  type    = "A"
  content = "192.168.0.50"
  ttl     = 3600
  proxied = false
  
  # Don't specify tags at all - this should NOT cause drift
}