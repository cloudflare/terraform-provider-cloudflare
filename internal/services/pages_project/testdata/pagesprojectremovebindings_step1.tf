resource "cloudflare_workers_kv_namespace" "%[1]s_kv" {
	account_id = "%[2]s"
	title = "tfacctest-pages-bindings-kv"
}

resource "cloudflare_pages_project" "%[1]s" {
	account_id = "%[2]s"
	name = "%[3]s"
	production_branch = "main"
	deployment_configs = {
		preview = {
			compatibility_date = "2023-06-01"
			compatibility_flags = []
			kv_namespaces = {
				KV_BINDING_1 = { namespace_id = cloudflare_workers_kv_namespace.%[1]s_kv.id }
				KV_BINDING_2 = { namespace_id = cloudflare_workers_kv_namespace.%[1]s_kv.id }
			}
			r2_buckets = {
				R2_BINDING = { name = "test-bucket" }
			}
			d1_databases = {
				D1_BINDING = { id = "445e2955-951a-4358-a35b-a4d0c813f63" }
			}
		}
		production = {
			compatibility_date = "2023-06-01"
			compatibility_flags = []
			kv_namespaces = {
				KV_BINDING = { namespace_id = cloudflare_workers_kv_namespace.%[1]s_kv.id }
			}
			r2_buckets = {
				R2_BINDING_1 = { name = "bucket-one" }
				R2_BINDING_2 = { name = "bucket-two" }
			}
		}
	}
}

