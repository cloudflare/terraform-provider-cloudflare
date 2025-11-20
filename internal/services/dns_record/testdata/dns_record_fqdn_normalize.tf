resource "cloudflare_dns_record" "%[1]s_subdomain" {
  zone_id = "%[2]s"
  name    = "test-fqdn"  # Subdomain without zone suffix
  content = "192.168.0.100"
  type    = "A"
  proxied = false
  ttl     = 3600
}

resource "cloudflare_dns_record" "%[1]s_subdomain_multi" {
  zone_id = "%[2]s"
  name    = "api.gateway.test"  # Multi-level subdomain
  content = "192.168.0.102"
  type    = "A"
  proxied = false
  ttl     = 3600
}
