resource "cloudflare_zero_trust_dlp_entry" "%[1]s" {
  account_id = "%[2]s"
  profile_id = "%[3]s"
  name       = "test-dlp-entry-validation-%[1]s"
  enabled    = true
  type       = "custom"
  
  pattern = {
    regex      = "4[0-9]{3}-[0-9]{4}-[0-9]{4}-[0-9]{4}"
    validation = "luhn"
  }
}
