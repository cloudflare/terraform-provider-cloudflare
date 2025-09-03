resource "cloudflare_zero_trust_dlp_entry" "%[1]s" {
  account_id = "%[2]s"
  profile_id = "%[3]s"
  name       = "test-dlp-entry-no-validation-%[1]s"
  enabled    = false
  type       = "custom"
  
  pattern = {
    regex = "\\b[A-Z0-9._%+-]+@[A-Z0-9.-]+\\.[A-Z]{2,}\\b"
  }
}
