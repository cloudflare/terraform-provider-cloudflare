resource "cloudflare_workflow" "%[1]s" {
  account_id = "%[2]s"
  # Missing required workflow_name, script_name, class_name
}