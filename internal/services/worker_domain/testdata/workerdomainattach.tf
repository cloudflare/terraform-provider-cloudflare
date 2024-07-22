
resource "cloudflare_worker_script" "%[1]s-script" {
  account_id = "%[3]s"
  name = "%[1]s"
  content = "%[2]s"
}

resource "cloudflare_worker_domain" "%[1]s" {
	zone_id = "%[5]s"
	account_id = "%[3]s"
	hostname = "%[4]s"
	service = "%[1]s"
	depends_on = [cloudflare_worker_script.%[1]s-script]
}