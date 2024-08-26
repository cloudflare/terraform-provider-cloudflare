
resource "cloudflare_workers_kv_namespace" "%[1]s" {
	account_id = "%[2]s"
	title = "%[1]s"
}