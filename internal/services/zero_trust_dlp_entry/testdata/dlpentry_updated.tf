resource "cloudflare_zero_trust_dlp_entry" "%[1]s" {
  account_id = "%[2]s"
  profile_id = "%[3]s"
  name       = "test-dlp-entry-%[1]s-updated"
  enabled    = false
  type       = "custom"
  
  pattern = {
    regex = "[0-9]{3}[[:space:]]?-?[0-9]{2}[[:space:]]?-?[0-9]{4}"
  }
}