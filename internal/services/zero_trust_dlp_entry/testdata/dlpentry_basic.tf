resource "cloudflare_zero_trust_dlp_entry" "%[1]s" {
  account_id = "%[2]s"
  profile_id = "%[3]s"
  name       = "test-dlp-entry-%[1]s"
  enabled    = true
  type       = "custom"
  
  pattern = {
    regex      = "[0-9]{4}[[:space:]]?-?[0-9]{4}[[:space:]]?-?[0-9]{4}[[:space:]]?-?[0-9]{4}"
    validation = "luhn"
  }
}