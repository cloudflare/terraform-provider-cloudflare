resource "cloudflare_zone_setting" "%[1]s" {
  zone_id = "%[2]s"
  setting_id = "nel"
  value = {
    enabled = true
  }
}
