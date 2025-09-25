resource "cloudflare_workflow" "%[1]s" {
  account_id    = "%[2]s"
  workflow_name = "%[3]s"
  script_name   = "%[4]s"
  class_name    = "%[5]s"
}