resource "cloudflare_zero_trust_dlp_dataset" "example_zero_trust_dlp_dataset" {
  account_id = "account_id"
  name = "name"
  case_sensitive = true
  description = "description"
  encoding_version = 0
  secret = true
}
