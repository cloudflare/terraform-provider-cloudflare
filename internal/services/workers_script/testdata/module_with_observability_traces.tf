resource "cloudflare_workers_script" "%[1]s" {
  account_id  = "%[2]s"
  script_name = "%[1]s"
  content     = "%[3]s"
  main_module = "worker.mjs"

  observability = {
    enabled = true
    traces = {
      enabled            = true
      head_sampling_rate = 1
      persist            = true
    }
  }
}
