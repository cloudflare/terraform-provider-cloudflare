resource "cloudflare_zone_setting" "%[1]s_zero_rtt" {
  zone_id    = "%[2]s"
  setting_id = "0rtt"
  value      = "on"
}
