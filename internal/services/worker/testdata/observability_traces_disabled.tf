resource "cloudflare_worker" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"

  observability = {
    enabled            = true
    head_sampling_rate = 1
    traces = {
      enabled = false
    }
  }
}
