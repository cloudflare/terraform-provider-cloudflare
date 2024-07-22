
resource "cloudflare_record" "%[3]s" {
	zone_id = "%[1]s"
	name = "%[2]s"
	value = "192.168.0.11"
	type = "A"
	ttl = 3600
	tags = ["updated_tag1", "updated_tag2"]
    comment = "this is am updated comment"
}