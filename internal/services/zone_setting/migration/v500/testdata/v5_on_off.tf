resource "cloudflare_zone_setting" "%[1]s_http3" {
  zone_id    = "%[2]s"
  setting_id = "http3"
  value      = "on"
}


removed {
  from = cloudflare_zone_settings_override.%[1]s
  lifecycle {
    destroy = false
  }
}
