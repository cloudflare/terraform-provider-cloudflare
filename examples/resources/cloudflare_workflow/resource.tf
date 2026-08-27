resource "cloudflare_workflow" "example_workflow" {
  account_id = "account_id"
  workflow_name = "x"
  class_name = "x"
  script_name = "x"
  concurrency = {
    limit = 1
  }
  default_retention = {
    error_retention = "5 minutes"
    success_retention = "5 minutes"
  }
  limits = {
    steps = 1
  }
  schedules = [{
    cron = "x"
  }]
}
