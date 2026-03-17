resource "cloudflare_zone_settings_override" {
  zone_id = "%[2]s"
  settings {
    http3             = "on"
    browser_cache_ttl = 14400
    min_tls_version   = "1.2"
  }
}
