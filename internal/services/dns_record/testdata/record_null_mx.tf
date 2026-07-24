
	resource "cloudflare_dns_record" "%[1]s" {
		zone_id  = "%[2]s"
		type     = "MX"
		name     = "%[1]s.%[3]s"
		data = {
			priority = 0
			target = "."
		}
		ttl = 300
	  }
