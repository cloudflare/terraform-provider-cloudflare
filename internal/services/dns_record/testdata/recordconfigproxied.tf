
resource "cloudflare_dns_record" "%[4]s" {
	zone_id = "%[1]s"
	name = "%[3]s"
	content = "%[2]s"
	type = "CNAME"
	proxied = true
	ttl = 1
}
