
resource "cloudflare_zero_trust_access_custom_page" "%[1]s" {
	account_id = "%[2]s"
	name = "%[1]s"
	type = "%[3]s"
	custom_html = "%[4]s"
}
	