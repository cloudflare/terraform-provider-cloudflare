
  resource "cloudflare_notification_policy_webhooks" "%[1]s" {
	account_id  = "%[2]s"
    name        = "my webhooks destination for receiving Cloudflare notifications"
    url         = "https://example.com"
    secret      =  "my-secret-key"
  }