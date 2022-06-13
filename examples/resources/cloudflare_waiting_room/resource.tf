# Waiting Room
resource "cloudflare_waiting_room" "example" {
  zone_id              = "ae36f999674d196762efcc5abb06b345"
  name                 = "foo"
  host                 = "foo.example.com"
  path                 = "/"
  new_users_per_minute = 200
  total_active_users   = 200
}