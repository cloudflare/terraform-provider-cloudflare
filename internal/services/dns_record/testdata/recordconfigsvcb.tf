
resource "cloudflare_dns_record" "%[2]s" {
	zone_id = "%[1]s"
	name = "%[2]s"
	type = "SVCB"
	data = {
  priority = "2"
		target   = "foo."
		value    = "alpn=\"h3,h2\""
}
	ttl = 300
}
