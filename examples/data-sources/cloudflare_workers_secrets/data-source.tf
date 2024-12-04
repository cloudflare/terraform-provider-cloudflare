data "cloudflare_workers_secrets" "example_workers_secrets" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  dispatch_namespace = "my-dispatch-namespace"
  script_name = "this-is_my_script-01"
}
