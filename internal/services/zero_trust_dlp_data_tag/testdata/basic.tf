resource "cloudflare_zero_trust_dlp_data_tag_category" "%[1]s" {
  account_id  = "%[2]s"
  name        = "tf-acc-%[1]s-category"
  description = "Parent category for data_tag acceptance test"
}

resource "cloudflare_zero_trust_dlp_data_tag" "%[1]s" {
  account_id  = "%[2]s"
  category_id = cloudflare_zero_trust_dlp_data_tag_category.%[1]s.id
  name        = "tf-acc-%[1]s"
  description = "Acceptance test data tag"
}
