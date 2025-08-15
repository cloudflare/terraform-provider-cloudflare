resource "cloudflare_bot_management" "%[1]s" {
	zone_id = "%[2]s"

	suppress_session_score = %[3]t
}