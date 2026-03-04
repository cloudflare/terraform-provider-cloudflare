resource "cloudflare_device_dex_test" "%s" {
  account_id  = "%s"
  name        = "%s"
  description = "Comprehensive test with all fields"
  interval    = "0h30m0s"
  enabled     = true

  data {
    kind   = "http"
    host   = "https://dash.cloudflare.com/login"
    method = "GET"
  }
}
