resource "cloudflare_device_dex_test" "example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  name        = "GET homepage"
  description = "Send a HTTP GET request to the home endpoint every half hour."
  interval    = "0h30m0s"
  enabled     = true
  data {
    host   = "https://example.com/home"
    kind   = "http"
    method = "GET"
  }
}
