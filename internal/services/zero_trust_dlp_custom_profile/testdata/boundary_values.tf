resource "cloudflare_zero_trust_dlp_custom_profile" "%[1]s" {
	account_id           = "%[2]s"
	name                 = "%[1]s-boundary"
	description          = "Testing boundary values"
	allowed_match_count  = 0
	confidence_threshold = "low"
	ocr_enabled          = false
	ai_context_enabled   = false
}