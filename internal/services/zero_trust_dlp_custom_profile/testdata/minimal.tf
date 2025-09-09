resource "cloudflare_zero_trust_dlp_custom_profile" "%[1]s" {
	account_id = "%[2]s"
	name       = "%[1]s"
}