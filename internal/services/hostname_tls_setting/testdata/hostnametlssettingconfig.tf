
resource "cloudflare_hostname_tls_setting" "%[5]s" {
	zone_id	= "%[1]s"
	hostname = "%[2]s"
	setting_id = "%[3]s"
	value = "%[4]s"
}
