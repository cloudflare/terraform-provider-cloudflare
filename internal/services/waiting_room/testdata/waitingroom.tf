
	resource "cloudflare_record" "%[1]s-shop-1" {
		zone_id = "%[3]s"
		name = "shop1"
		value = "192.168.0.10"
		type = "A"
		ttl = 3600
	}

	resource "cloudflare_record" "%[1]s-shop-2" {
		zone_id = "%[3]s"
		name = "shop2"
		value = "192.168.0.11"
		type = "A"
		ttl = 3600
	}

resource "cloudflare_waiting_room" "%[1]s" {
  name                      = "%[2]s"
  zone_id                   = "%[3]s"
  host                      = "%[4]s"
  new_users_per_minute      = 400
  total_active_users        = 405
  path                      = "%[5]s"
  session_duration          = 10
  queueing_method           = "fifo"
  custom_page_html          = "foobar"
  default_template_language = "en-US"
  description               = "my desc"
  disable_session_renewal   = true
  suspended                 = true
  queue_all                 = false
  json_response_enabled     = true
  cookie_suffix             = "queue1"
  additional_routes =[ {
    host = "shop1.%[4]s"
    path = "%[5]s"
  },
    {
    host = "shop2.%[4]s"
    }]


  queueing_status_code      = 200

  depends_on = [cloudflare_record.%[1]s-shop-1, cloudflare_record.%[1]s-shop-2]
}
