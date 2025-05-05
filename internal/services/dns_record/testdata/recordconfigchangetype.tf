
resource "cloudflare_dns_record" "%[4]s" {
	zone_id = "%[1]s"
	name = "%[2]s.%[3]s"
	content = "%[3]s"
	type = "CNAME"
	ttl = 3600
}
