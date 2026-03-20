resource "cloudflare_zone_setting" "%[1]s_browser_cache_ttl" {
  zone_id    = "%[2]s"
  setting_id = "browser_cache_ttl"
  value      = 14400
}
