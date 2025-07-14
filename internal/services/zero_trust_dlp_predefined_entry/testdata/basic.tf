resource "cloudflare_zero_trust_dlp_predefined_entry" "%[1]s" {
  account_id = "%[2]s"
  entry_id = "%[3]s"
  enabled = "%[4]t"
}
