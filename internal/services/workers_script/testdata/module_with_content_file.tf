resource "cloudflare_workers_script" "%[1]s" {
  account_id = "%[2]s"
  script_name = "%[1]s"
  content_file = "%[3]s"
  content_sha256 = filesha256("%[3]s")
  main_module = "worker.mjs"
}
