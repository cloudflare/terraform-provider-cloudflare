
	resource "cloudflare_zero_trust_access_application" "%[1]s" {
		zone_id = "%[2]s"
		name = "%[1]s"
		domain = "%[1]s.%[3]s"
	}

	data "cloudflare_zero_trust_access_application" "%[1]s" {
		zone_id = "%[2]s"
		domain = "%[1]s.%[3]s"
		depends_on = [cloudflare_zero_trust_access_application.%[1]s]
	}
	