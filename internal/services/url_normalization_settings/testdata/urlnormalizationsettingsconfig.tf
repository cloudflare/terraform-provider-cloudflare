
resource "cloudflare_url_normalization_settings" "%[4]s" {
	zone_id = "%[1]s"
	type = "%[2]s"
	scope = "%[3]s"
}
