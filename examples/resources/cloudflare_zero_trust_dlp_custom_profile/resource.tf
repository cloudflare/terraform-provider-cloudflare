resource "cloudflare_zero_trust_dlp_custom_profile" "example_zero_trust_dlp_custom_profile" {
  account_id = "account_id"
  name = "name"
  ai_context_enabled = true
  allowed_match_count = 5
  confidence_threshold = "confidence_threshold"
  context_awareness = {
    enabled = true
    skip = {
      files = true
    }
  }
  data_classes = ["182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"]
  data_tags = ["182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"]
  description = "description"
  ocr_enabled = true
  sensitivity_levels = [{
    group_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
    level_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
  }]
  shared_entries = [{
    enabled = true
    entry_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
  }]
}
