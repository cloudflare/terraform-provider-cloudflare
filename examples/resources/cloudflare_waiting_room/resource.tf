# Waiting Room
resource "cloudflare_waiting_room" "example" {
  zone_id              = "0da42c8d2132a9ddaf714f9e7c920711"
  name                 = "foo"
  host                 = "foo.example.com"
  path                 = "/"
  new_users_per_minute = 200
  total_active_users   = 200
}