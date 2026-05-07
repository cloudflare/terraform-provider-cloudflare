resource "cloudflare_zero_trust_dlp_settings" "example_zero_trust_dlp_settings" {
  account_id = "account_id"
  ai_context_analysis = true
  ocr = true
  payload_logging = {
    masking_level = "full"
    public_key = "public_key"
  }
}
