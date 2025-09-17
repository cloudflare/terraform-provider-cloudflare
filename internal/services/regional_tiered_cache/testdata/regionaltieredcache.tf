
resource "cloudflare_regional_tiered_cache" "%[1]s" {
  zone_id = "%[2]s"
  value   = "%[3]s"
}
