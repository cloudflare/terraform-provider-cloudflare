resource "cloudflare_zero_trust_dlp_custom_entry" "%[1]s" {
  account_id = "%[2]s"
  profile_id = "%[3]s"

  enabled = "%[4]b"
}
