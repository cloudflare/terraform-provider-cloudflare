resource "cloudflare_waiting_room" "%[1]s" {
  name                 = "%[2]s"
  zone_id              = "%[3]s"
  host                 = "%[4]s"
  new_users_per_minute = 200
  total_active_users   = 200
  session_duration     = 1
}
