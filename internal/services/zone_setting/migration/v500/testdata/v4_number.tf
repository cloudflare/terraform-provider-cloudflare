resource "cloudflare_zone_settings_override" {
  zone_id = "%[2]s"
  settings {
    browser_cache_ttl = 14400
  }
}
