resource "cloudflare_zero_trust_dlp_predefined_profile" "example_zero_trust_dlp_predefined_profile" {
  account_id = "account_id"
  profile_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
  ai_context_enabled = true
  allowed_match_count = 5
  confidence_threshold = "confidence_threshold"
  enabled_entries = ["182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"]
  entries = [{
    id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
    enabled = true
  }]
  ocr_enabled = true
}
