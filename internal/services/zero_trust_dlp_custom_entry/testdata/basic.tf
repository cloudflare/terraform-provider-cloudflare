resource "cloudflare_zero_trust_dlp_custom_entry" "%[1]s" {
  name = "%[1]s"
  account_id = "%[2]s"
  profile_id = "%[3]s"
  pattern = {
    regex = "customentryregex"
  }

  enabled = %[4]s
}
