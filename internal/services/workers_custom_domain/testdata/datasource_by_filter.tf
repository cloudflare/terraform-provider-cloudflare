resource "cloudflare_workers_custom_domain" "%[1]s" {
	zone_id = "%[4]s"
	account_id = "%[2]s"
	hostname = "%[3]s"
	service = "mute-truth-fdb1"
	environment = "production"
}

data "cloudflare_workers_custom_domain" "%[1]s" {
	account_id = "%[2]s"
	filter = {
		hostname = cloudflare_workers_custom_domain.%[1]s.hostname
		service = "mute-truth-fdb1"
	}
}
