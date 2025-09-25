resource "cloudflare_zero_trust_dlp_entry" "%[1]s" {
  account_id = "%[2]s"
  profile_id = "%[3]s"
  name       = "toggle-enabled-%[1]s"
  enabled    = false
  type       = "custom"
  
  pattern = {
    regex = "\\b[0-9]{4}\\b"
  }
}
