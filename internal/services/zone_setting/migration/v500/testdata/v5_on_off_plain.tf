resource "cloudflare_zone_setting" {
  zone_id    = "%[2]s"
  setting_id = "http3"
  value      = "on"
}
