# Waiting Room Event
resource "cloudflare_waiting_room_event" "example" {
  zone_id              = "0da42c8d2132a9ddaf714f9e7c920711"
  waiting_room_id      = "d41d8cd98f00b204e9800998ecf8427e"
  name                 = "foo"
  event_start_time     = "2006-01-02T15:04:05Z"
  event_end_time       = "2006-01-02T20:04:05Z"
}