resource "cloudflare_workers_kv_namespace" "example_ns" {
  title = "test-namespace"
}

resource "cloudflare_workers_kv" "example" {
  namespace_id = cloudflare_workers_kv_namespace.example_ns.id
  key          = "test-key"
  value        = "test value"
}
