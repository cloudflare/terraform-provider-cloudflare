resource "cloudflare_zero_trust_dlp_custom_profile" "%[1]s" {
	account_id           = "%[2]s"
	name                 = "%[1]s-max"
	description          = "Profile with all optional attributes set"
	allowed_match_count  = 1000
	confidence_threshold = "high"
	ocr_enabled          = true
	ai_context_enabled   = true
}