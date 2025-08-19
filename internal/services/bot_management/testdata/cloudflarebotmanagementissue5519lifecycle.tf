resource "cloudflare_bot_management" "%[1]s" {
	zone_id = "%[2]s"

	enable_js = true

	lifecycle {
		ignore_changes = [
			suppress_session_score,
			ai_bots_protection,
			crawler_protection,
			enable_js
		]
	}
}