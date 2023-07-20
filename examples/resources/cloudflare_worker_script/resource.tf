resource "cloudflare_workers_kv_namespace" "my_namespace" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  title      = "example"
}

# Sets the script with the name "script_1"
resource "cloudflare_worker_script" "my_script" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "script_1"
  content    = file("script.js")

  kv_namespace_binding {
    name         = "MY_EXAMPLE_KV_NAMESPACE"
    namespace_id = cloudflare_workers_kv_namespace.my_namespace.id
  }

  plain_text_binding {
    name = "MY_EXAMPLE_PLAIN_TEXT"
    text = "foobar"
  }

  secret_text_binding {
    name = "MY_EXAMPLE_SECRET_TEXT"
    text = var.secret_foo_value
  }

  webassembly_binding {
    name   = "MY_EXAMPLE_WASM"
    module = filebase64("example.wasm")
  }

  service_binding {
    name        = "MY_SERVICE_BINDING"
    service     = "MY_SERVICE"
    environment = "production"
  }

  r2_bucket_binding {
    name        = "MY_BUCKET"
    bucket_name = "MY_BUCKET_NAME"
  }

  analytics_engine_binding {
    name    = "MY_DATASET"
    dataset = "dataset1"
  }
}
