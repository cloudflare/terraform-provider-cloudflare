
	resource "cloudflare_bot_management" "%[1]s" {
		zone_id = "%[2]s"

		enable_js = "%[3]t"

		suppress_session_score = "%[4]t"
		auto_update_model = "%[5]t"
	}
