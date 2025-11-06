resource "cloudflare_universal_ssl_setting" "%[1]s" {
	zone_id = "%[2]s"
	enabled = true
}
