
resource "cloudflare_workers_script" "%[1]s" {
  account_id = "%[3]s"
  script_name = "%[1]s"
  content = "%[2]s"
}
