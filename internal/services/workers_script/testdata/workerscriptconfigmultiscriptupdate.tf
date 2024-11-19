
resource "cloudflare_workers_script" "%[1]s" {
  account_id = "%[3]s"
  name = "%[1]s"
  content = "%[2]s"
}