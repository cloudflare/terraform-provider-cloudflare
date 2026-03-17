resource "cloudflare_zone_setting" "%[1]s_browser_cache_ttl" {
  zone_id    = "%[2]s"
  setting_id = "browser_cache_ttl"
  value      = 14400
}

import {
  to = cloudflare_zone_setting.%[1]s_browser_cache_ttl
  id = "%[2]s/browser_cache_ttl"
}

removed {
  from = cloudflare_zone_settings_override.%[1]s
  lifecycle {
    destroy = false
  }
}
