
resource "cloudflare_record" "%[3]s" {
	zone_id = "%[1]s"
	name = "%[2]s"
	value = "192.168.0.10"
	type = "A"
	ttl = 3600
}