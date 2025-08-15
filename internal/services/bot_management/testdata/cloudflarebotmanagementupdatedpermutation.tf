resource "cloudflare_bot_management" "%[1]s" {
	zone_id = "%[2]s"

	enable_js = false
	suppress_session_score = true
	ai_bots_protection = "disabled"
}