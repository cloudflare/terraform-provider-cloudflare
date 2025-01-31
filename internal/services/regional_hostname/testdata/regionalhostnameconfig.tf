
resource "cloudflare_regional_hostname" "%[1]s" {
	zone_id = "%[2]s"
	hostname = "%[3]s"
	region_key = "%[4]s"
}