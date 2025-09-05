resource "cloudflare_zero_trust_dlp_custom_profile" "%[1]s" {
	account_id  = "%[2]s"
	name        = "%[1]s"
	description = "Test custom DLP profile"

	allowed_match_count = 5
	ocr_enabled         = true
	ai_context_enabled  = false
}
