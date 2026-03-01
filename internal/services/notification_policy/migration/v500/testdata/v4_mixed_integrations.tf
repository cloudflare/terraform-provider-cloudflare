resource "cloudflare_notification_policy_webhooks" "test_webhook" {
  account_id = "%s"
  name       = "tf-test-webhook-%s"
  url        = "https://www.cloudflare.com/cdn-cgi/trace"
}

resource "cloudflare_notification_policy" "%s" {
  account_id  = "%s"
  name        = "%s"
  description = "Test notification policy with mixed integrations"
  enabled     = true
  alert_type  = "expiring_service_token_alert"

  email_integration {
    id = "mixed-test1@example.com"
  }

  email_integration {
    id = "mixed-test2@example.com"
  }

  webhooks_integration {
    id   = cloudflare_notification_policy_webhooks.test_webhook.id
    name = cloudflare_notification_policy_webhooks.test_webhook.name
  }
}
