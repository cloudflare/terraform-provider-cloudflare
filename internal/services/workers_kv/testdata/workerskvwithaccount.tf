
resource "cloudflare_workers_kv" "%[1]s" {
	account_id = "%[4]s"
	namespace_id = cloudflare_workers_kv_namespace.%[1]s.id
	key = "%[2]s"
	value = "%[3]s"
}