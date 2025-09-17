
resource "cloudflare_hostname_tls_setting" "%[4]s" {
	zone_id	= "%[1]s"
	hostname = "%[2]s"
	setting_id = "%[3]s"
	value = [
    "AES128-SHA",
    "ECDHE-RSA-AES256-SHA",
  ]
}
