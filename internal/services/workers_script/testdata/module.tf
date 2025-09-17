resource "cloudflare_workers_kv_namespace" "%[1]s" {
	account_id = "%[3]s"
	title = "%[1]s"
}

resource "cloudflare_queue" "%[1]s" {
	account_id = "%[3]s"
	queue_name = "%[1]s"
}

resource "cloudflare_workers_script" "%[1]s" {
  account_id = "%[3]s"
  script_name = "%[1]s"
  content = "%[2]s"
  main_module = "worker.mjs"
  compatibility_date = "%[4]s"
  compatibility_flags = ["%[5]s"]
  placement = {
    mode = "smart"
  }
  migrations = {
    new_tag = "v1"
    new_sqlite_classes = ["MyDurableObject"]
  }
  bindings = [
    {
      name = "MY_KV_NAMESPACE"
      type = "kv_namespace"
      namespace_id = cloudflare_workers_kv_namespace.%[1]s.id
    },
    {
      name = "MY_QUEUE"
      type = "queue"
      queue_name = cloudflare_queue.%[1]s.queue_name
    },
    {
      name = "SECRET"
      type = "secret_text"
      text = "shhh!!"
    }
  ]
  observability = {
    enabled = true
    head_sampling_rate = 0.1
  }
}
