resource "cloudflare_worker" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"

  observability = {
    enabled            = false
    head_sampling_rate = 1
    logs = {
      enabled            = false
      head_sampling_rate = 1
    }
  }
}
