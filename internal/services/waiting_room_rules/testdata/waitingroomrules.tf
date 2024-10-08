
resource "cloudflare_waiting_room" "%[1]s" {
  name                      = "%[4]s"
  zone_id                   = "%[2]s"
  host                      = "www.%[3]s"
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

resource "cloudflare_waiting_room_rules" "%[1]s" {
  zone_id		  = "%[2]s"
  waiting_room_id = cloudflare_waiting_room.%[1]s.id

  rules {
    action      = "bypass_waiting_room"
    expression  = "ip.src in {192.0.2.1}"
    description = "ip bypass"
    status 	    = "enabled"
  }

  rules {
    action      = "bypass_waiting_room"
    expression 	= "http.request.uri.query contains \"bypass=true\""
    description = "query string bypass"
    status 	    = "disabled"
  }
}
