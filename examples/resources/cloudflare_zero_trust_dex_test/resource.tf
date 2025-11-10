resource "cloudflare_zero_trust_dex_test" "example_zero_trust_dex_test" {
  account_id = "01a7362d577a6c3019a474fd6f485823"
  data = {
    host = "https://dash.cloudflare.com"
    kind = "http"
    method = "GET"
  }
  enabled = true
  interval = "30m"
  name = "HTTP dash health check"
  description = "Checks the dash endpoint every 30 minutes"
  target_policies = [{
    id = "id"
    default = true
    name = "name"
  }]
  targeted = true
}
