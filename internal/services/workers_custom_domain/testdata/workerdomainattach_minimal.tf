resource "cloudflare_workers_custom_domain" "%[1]s" {
	account_id = "%[2]s"
	hostname = "%[3]s"
	service = "mute-truth-fdb1"
}
