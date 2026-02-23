resource "cloudflare_device_dex_test" "%s" {
  account_id  = "%s"
  name        = "%s"
  description = "Test network path to Cloudflare DNS"
  interval    = "1h0m0s"
  enabled     = true

  data {
    kind = "traceroute"
    host = "1.1.1.1"
  }
}
