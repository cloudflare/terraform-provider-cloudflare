resource "cloudflare_workers_deployment" "example_workers_deployment" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  script_name = "this-is_my_script-01"
  strategy = "percentage"
  versions = [{
    percentage = 100
    version_id = "023e105f-2a42-4f8b-a1c1-73f6a2a30c0f"
  }]
  annotations = {
    workers_message = "Deploy bug fix."
  }
}
