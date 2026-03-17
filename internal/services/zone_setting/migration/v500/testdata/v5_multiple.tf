resource "cloudflare_zone_setting" "%[1]s_browser_cache_ttl" {
  zone_id    = "%[2]s"
  setting_id = "browser_cache_ttl"
  value      = 14400
}

import {
  to = cloudflare_zone_setting.%[1]s_browser_cache_ttl
  id = "%[2]s/browser_cache_ttl"
}

resource "cloudflare_zone_setting" "%[1]s_http3" {
  zone_id    = "%[2]s"
  setting_id = "http3"
  value      = "on"
}

import {
  to = cloudflare_zone_setting.%[1]s_http3
  id = "%[2]s/http3"
}

resource "cloudflare_zone_setting" "%[1]s_min_tls_version" {
  zone_id    = "%[2]s"
  setting_id = "min_tls_version"
  value      = "1.2"
}

import {
  to = cloudflare_zone_setting.%[1]s_min_tls_version
  id = "%[2]s/min_tls_version"
}

removed {
  from = cloudflare_zone_settings_override.%[1]s
  lifecycle {
    destroy = false
  }
}
