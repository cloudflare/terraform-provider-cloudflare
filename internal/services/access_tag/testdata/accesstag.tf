
resource "cloudflare_access_tag" "%[1]s" {
	zone_id = "%[2]s"
	name = "%[1]s"
}
	