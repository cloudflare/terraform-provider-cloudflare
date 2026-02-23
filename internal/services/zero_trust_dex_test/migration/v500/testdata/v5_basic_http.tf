resource "cloudflare_zero_trust_dex_test" "%s" {
  account_id  = "%s"
  name        = "%s"
  description = "Test HTTP connectivity"
  interval    = "0h30m0s"
  enabled     = true

  data = {
    kind   = "http"
    host   = "https://dash.cloudflare.com"
    method = "GET"
  }
}
