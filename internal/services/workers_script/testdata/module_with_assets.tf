resource "cloudflare_workers_script" "%[1]s" {
  account_id = "%[2]s"
  script_name = "%[1]s"
  assets = {
    directory = "%[3]s"
  }
}
