resource "cloudflare_workers_kv_namespace" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  title      = "test-namespace"
}
