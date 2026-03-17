resource "cloudflare_zone_setting" {
  zone_id    = "%[2]s"
  setting_id = "browser_cache_ttl"
  value      = 14400
}
