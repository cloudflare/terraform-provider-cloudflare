resource "cloudflare_zero_trust_casb_webhook" "example_zero_trust_casb_webhook" {
  account_id = "46148281d8a93d002ef242d8b0d5f9f6"
  authentication_type = "Bearer Auth"
  destination_url = "https://example.com/webhook"
  label = "Send to Slack"
  headers = [{
    key = "Authorization"
    value = "Bearer token123"
  }, {
    key = "X-Custom-Header"
    value = "value"
  }]
  signing_secret = "my-secret-key"
  status = "enabled"
}
