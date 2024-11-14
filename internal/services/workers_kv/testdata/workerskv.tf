resource "cloudflare_workers_kv_namespace" "%[1]s" {
	account_id = "%[4]s"
	title = "%[1]s"
}

resource "cloudflare_workers_kv" "%[1]s" {
	account_id = "%[4]s"
	namespace_id = cloudflare_workers_kv_namespace.%[1]s.id
	key_name = "%[2]s"
	value = "%[3]s"
}
