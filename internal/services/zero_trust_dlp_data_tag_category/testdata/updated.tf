resource "cloudflare_zero_trust_dlp_data_tag_category" "%[1]s" {
  account_id  = "%[2]s"
  name        = "tf-acc-%[1]s-renamed"
  description = "Updated description"
}
