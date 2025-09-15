resource "cloudflare_workers_script" "%[1]s" {
  account_id   = "%[2]s"
  script_name  = "%[1]s"
  content      = "%[3]s"

  assets {
    config {
      run_worker_first = %[4]s
    }
  }
}
