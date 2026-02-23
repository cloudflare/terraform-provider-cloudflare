resource "cloudflare_device_dex_test" "%s" {
  account_id  = "%s"
  name        = "%s"
  description = "Minimal test"
  interval    = "0h30m0s"
  enabled     = true

  data {
    kind   = "http"
    host   = "https://cloudflare.com"
    method = "GET"
  }
}
