resource "cloudflare_workers_for_platforms_namespace" "my_namespace" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "example-namespace"
}

resource "cloudflare_worker_script" "customer_worker_1" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "customer-worker-1"
  content    = file("script.js")
  dispatch_namespace = "example-namespace"
  tags = ["free"]
}
