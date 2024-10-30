
resource "cloudflare_workers_kv_namespace" "%[1]s" {
	account_id = "%[4]s"
	title = "%[1]s"
}

resource "cloudflare_queue" "%[1]s" {
	account_id = "%[4]s"
	name = "%[1]s"
}

resource "cloudflare_workers_script" "%[1]s-service" {
	account_id = "%[4]s"
	name    = "%[1]s-service"
	content = "%[2]s"
}

resource "cloudflare_workers_script" "%[1]s" {
  account_id = "%[4]s"
  script_name    = "%[1]s"
  content = "%[2]s"

  kv_namespace_binding =[ {
    name         = "MY_KV_NAMESPACE"
    namespace_id = cloudflare_workers_kv_namespace.%[1]s.id
  }

  plain_text_binding {
    name = "MY_PLAIN_TEXT"
    text = "%[1]s"
  }]

  secret_text_binding =[ {
    name = "MY_SECRET_TEXT"
    text = "%[1]s"
  }]

  webassembly_binding =[ {
    name = "MY_WASM"
    module = "%[3]s"
  }]

  r2_bucket_binding =[ {
	name = "MY_BUCKET"
	bucket_name = "%[1]s"
  }]

  service_binding =[ {
	name = "MY_SERVICE_BINDING"
    service = cloudflare_workers_script.%[1]s-service.name
    environment = "production"
  }

  queue_binding =[ {
    binding         = "MY_QUEUE"
    queue = cloudflare_queue.%[1]s.name
  }

}]
