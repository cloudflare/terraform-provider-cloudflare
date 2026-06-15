resource "cloudflare_zero_trust_dlp_sensitivity_level_order" "example_zero_trust_dlp_sensitivity_level_order" {
  account_id = "account_id"
  sensitivity_group_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
  level_ids = ["182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"]
}
