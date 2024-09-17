
resource "cloudflare_record" "%[2]s" {
	zone_id = "%[1]s"
	name = "@"
	value = "192.168.0.10"
	type = "A"
	ttl = 3600
}