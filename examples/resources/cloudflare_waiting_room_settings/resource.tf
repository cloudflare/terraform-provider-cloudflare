resource "cloudflare_waiting_room_settings" "example" {
  zone_id                      = "0da42c8d2132a9ddaf714f9e7c920711"
  search_engine_crawler_bypass = true
}
