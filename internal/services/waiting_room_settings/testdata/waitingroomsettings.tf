
resource "cloudflare_waiting_room_settings" "%[1]s" {
  zone_id                      = "%[2]s"
  search_engine_crawler_bypass = true
}
