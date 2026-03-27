resource "cloudflare_zone_setting" "%[1]s_security_header" {
  zone_id    = "%[2]s"
  setting_id = "security_header"
  value = {
    strict_transport_security = {
      enabled            = true
      include_subdomains = true
      max_age            = 86400
      nosniff            = false
      preload            = false
    }
  }
}

import {
  to = cloudflare_zone_setting.%[1]s_security_header
  id = "%[2]s/security_header"
}

removed {
  from = cloudflare_zone_settings_override.%[1]s
  lifecycle {
    destroy = false
  }
}
