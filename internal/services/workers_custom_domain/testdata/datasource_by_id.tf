resource "cloudflare_workers_custom_domain" "%[1]s" {
	zone_id = "%[4]s"
	account_id = "%[2]s"
	hostname = "%[3]s"
	service = "mute-truth-fdb1"
	environment = "production"
}

data "cloudflare_workers_custom_domain" "%[1]s" {
	account_id = "%[2]s"
	domain_id = cloudflare_workers_custom_domain.%[1]s.id
}
