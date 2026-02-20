resource "cloudflare_authenticated_origin_pulls_settings" "%[1]s" {
  zone_id = "%[2]s"
  enabled = true
}
