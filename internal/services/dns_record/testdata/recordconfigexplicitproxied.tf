
resource "cloudflare_dns_record" "%[2]s" {
	zone_id = "%[1]s"
	name = "%[2]s.%[3]s"
	content = "%[3]s"
	type = "CNAME"
	proxied = %[4]s
	ttl = %[5]s
}
