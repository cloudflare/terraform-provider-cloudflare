resource "cloudflare_zero_trust_dlp_data_class" "example_zero_trust_dlp_data_class" {
  account_id = "account_id"
  data_tags = ["182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"]
  expression = "expression"
  name = "name"
  sensitivity_levels = [{
    group_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
    level_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
  }]
  description = "description"
}
