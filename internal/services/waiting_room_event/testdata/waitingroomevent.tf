resource "cloudflare_waiting_room" "%[1]s" {
  name                      = "%[8]s"
  zone_id                   = "%[3]s"
  host                      = "www.%[7]s"
  new_users_per_minute      = 400
  total_active_users        = 405
  path                      = "/foobar"
  session_duration          = 10
  custom_page_html          = "foobar"
  description               = "my desc"
  disable_session_renewal   = true
  suspended                 = true
  queue_all                 = false
  json_response_enabled     = true
}

resource "cloudflare_waiting_room_event" "%[1]s" {
  name                      = "%[2]s"
  zone_id                   = "%[3]s"
  waiting_room_id           = cloudflare_waiting_room.%[1]s.id
  event_start_time          = "%[5]s"
  event_end_time            = "%[6]s"
  total_active_users        = 405
  new_users_per_minute      = 400
  custom_page_html          = "foobar"
  queueing_method           = "fifo"
  shuffle_at_event_start    = false
  disable_session_renewal   = true
  suspended                 = true
  description               = "my desc"
  session_duration          = 10
}
