resource "cloudflare_zero_trust_dlp_custom_profile" "%[1]s" {
	account_id  = "%[2]s"
	name        = "%[1]s-all-types"
	description = "Profile without shared entries for basic testing"
	
	allowed_match_count = 10
	ocr_enabled         = true
	ai_context_enabled  = true
}