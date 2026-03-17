resource "cloudflare_zone_settings_override" {
  zone_id = "%[2]s"
  settings {
    http3 = "on"
  }
}
