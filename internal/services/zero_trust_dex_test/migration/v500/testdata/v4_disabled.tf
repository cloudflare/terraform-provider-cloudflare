resource "cloudflare_device_dex_test" "%s" {
  account_id  = "%s"
  name        = "%s"
  description = "Currently disabled for maintenance"
  interval    = "0h15m0s"
  enabled     = false

  data {
    kind   = "http"
    host   = "https://internal.example.com"
    method = "GET"
  }
}
