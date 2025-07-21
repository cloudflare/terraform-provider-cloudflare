data "cloudflare_workers_deployment" "example_workers_deployment" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  script_name = "this-is_my_script-01"
  deployment_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
}
