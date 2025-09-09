resource "cloudflare_zero_trust_dlp_custom_profile" "%[1]s" {
	account_id  = "%[2]s"
	name        = "%[1]s-updated"
	description = "Updated test custom DLP profile"
	
	allowed_match_count = 10
	ocr_enabled         = false
	ai_context_enabled  = true
}