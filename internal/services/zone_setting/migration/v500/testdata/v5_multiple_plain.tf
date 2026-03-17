resource "cloudflare_zone_setting" "%[1]s_browser_cache_ttl" {
  zone_id    = "%[2]s"
  setting_id = "browser_cache_ttl"
  value      = 14400
}

resource "cloudflare_zone_setting" "%[1]s_http3" {
  zone_id    = "%[2]s"
  setting_id = "http3"
  value      = "on"
}

resource "cloudflare_zone_setting" "%[1]s_min_tls_version" {
  zone_id    = "%[2]s"
  setting_id = "min_tls_version"
  value      = "1.2"
}
