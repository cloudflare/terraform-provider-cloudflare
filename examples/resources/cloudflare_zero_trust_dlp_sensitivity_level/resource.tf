resource "cloudflare_zero_trust_dlp_sensitivity_level" "example_zero_trust_dlp_sensitivity_level" {
  account_id = "account_id"
  sensitivity_group_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
  name = "name"
  description = "description"
}
