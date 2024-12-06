data "cloudflare_workers_secret" "example_workers_secret" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  dispatch_namespace = "my-dispatch-namespace"
  script_name = "this-is_my_script-01"
  secret_name = "mySecret"
}
