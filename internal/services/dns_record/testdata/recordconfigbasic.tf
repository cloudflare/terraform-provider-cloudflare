
resource "cloudflare_dns_record" "%[3]s" {
	zone_id = "%[1]s"
	name = "%[2]s.%[3]s.%[4]s"
	content = "192.168.0.10"
	type = "A"
	ttl = 3600
	tags = ["tag1", "tag2"]
  comment = "this is a comment"
}
