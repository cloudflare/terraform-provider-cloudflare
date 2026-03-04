resource "cloudflare_notification_policy_webhooks" "webhook_1" {
  account_id = "%s"
  name       = "tf-test-webhook-1-%s"
  url        = "https://www.cloudflare.com/cdn-cgi/trace"
}

resource "cloudflare_notification_policy_webhooks" "webhook_2" {
  account_id = "%s"
  name       = "tf-test-webhook-2-%s"
  url        = "https://www.cloudflare.com/cdn-cgi/trace"
}

resource "cloudflare_notification_policy_webhooks" "webhook_3" {
  account_id = "%s"
  name       = "tf-test-webhook-3-%s"
  url        = "https://www.cloudflare.com/cdn-cgi/trace"
}

resource "cloudflare_notification_policy" "%s" {
  account_id  = "%s"
  name        = "%s"
  description = "Test notification policy with multiple webhooks"
  enabled     = true
  alert_type  = "universal_ssl_event_type"

  webhooks_integration {
    id   = cloudflare_notification_policy_webhooks.webhook_1.id
    name = cloudflare_notification_policy_webhooks.webhook_1.name
  }

  webhooks_integration {
    id   = cloudflare_notification_policy_webhooks.webhook_2.id
    name = cloudflare_notification_policy_webhooks.webhook_2.name
  }

  webhooks_integration {
    id   = cloudflare_notification_policy_webhooks.webhook_3.id
    name = cloudflare_notification_policy_webhooks.webhook_3.name
  }
}
