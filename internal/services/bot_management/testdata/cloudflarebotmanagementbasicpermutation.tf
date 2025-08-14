resource "cloudflare_bot_management" "%[1]s" {
	zone_id = "%[2]s"

	enable_js = true
	suppress_session_score = false
	ai_bots_protection = "block"
}