resource "cloudflare_zero_trust_dlp_predefined_profile" "%[1]s" {
  account_id          = "%[2]s"
  profile_id          = "%[3]s"
  allowed_match_count = 3
  enabled_entries     = ["%[4]s"]
}
