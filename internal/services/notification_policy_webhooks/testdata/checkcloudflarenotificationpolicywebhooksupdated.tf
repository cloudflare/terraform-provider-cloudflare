
  resource "cloudflare_notification_policy_webhooks" "%[1]s" {
	account_id  = "%[3]s"
    name        = "%[2]s"
	url         = "https://example.com"
  }