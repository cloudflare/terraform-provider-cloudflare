
resource "cloudflare_record" "%[4]s" {
	zone_id = "%[1]s"
	name = "%[3]s"
	value = "%[2]s"
	type = "CNAME"
	proxied = true
}