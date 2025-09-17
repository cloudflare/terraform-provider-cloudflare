
resource "cloudflare_workers_kv_namespace" "%[1]s" {
	account_id = "%[3]s"
	title = "%[2]s"
}
