resource "cloudflare_workers_kv_namespace" "example_ns" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  title      = "test-namespace"
}

resource "cloudflare_workers_kv" "example" {
  account_id   = "f037e56e89293a057740de681ac9abbe"
  namespace_id = cloudflare_workers_kv_namespace.example_ns.id
  key          = "test-key"
  value        = "test value"
}
