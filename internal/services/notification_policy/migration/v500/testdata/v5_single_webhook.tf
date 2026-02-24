resource "cloudflare_notification_policy_webhooks" "test_webhook" {
  account_id = "%s"
  name       = "tf-test-webhook-%s"
  url        = "https://www.cloudflare.com/cdn-cgi/trace"
}

resource "cloudflare_notification_policy" "%s" {
  account_id  = "%s"
  name        = "%s"
  description = "Test notification policy with single webhook"
  enabled     = true
  alert_type  = "universal_ssl_event_type"

  mechanisms = {
    webhooks = [{
      id = cloudflare_notification_policy_webhooks.test_webhook.id
    }]
  }
}
