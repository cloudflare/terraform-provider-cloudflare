data "cloudflare_workers_kv_namespaces" "example_workers_kv_namespaces" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  direction = "asc"
  order = "id"
}
