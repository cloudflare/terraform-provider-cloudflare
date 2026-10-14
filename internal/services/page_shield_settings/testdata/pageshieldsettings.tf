
resource "cloudflare_page_shield_settings" "%[1]s" {
  zone_id = "%[2]s"
  enabled = true
}
