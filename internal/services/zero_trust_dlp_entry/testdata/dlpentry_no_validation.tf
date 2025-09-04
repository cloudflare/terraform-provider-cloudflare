resource "cloudflare_zero_trust_dlp_entry" "%[1]s" {
  account_id = "%[2]s"
  profile_id = "%[3]s"
  name       = "test-dlp-entry-no-validation-%[1]s"
  enabled    = false
  type       = "custom"
  
  pattern = {
    regex = "\\b[a-z]{1,10}@[a-z]{1,10}\\.[a-z]{2,4}\\b"
  }
}
