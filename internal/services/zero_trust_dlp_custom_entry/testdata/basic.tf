resource "cloudflare_zero_trust_dlp_custom_profile" "custom_profile" {
  name = "%[1]s"
  account_id = "%[2]s"
}

resource "cloudflare_zero_trust_dlp_custom_entry" "%[1]s" {
  name = "%[1]s"
  account_id = "%[2]s"
  profile_id = cloudflare_zero_trust_dlp_custom_profile.custom_profile.id
  pattern = {
    regex = "customentryregex"
  }

  enabled = %[3]s
}
