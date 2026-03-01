resource "cloudflare_notification_policy" "%s" {
  account_id  = "%s"
  name        = "%s"
  description = "Test notification policy with multiple emails"
  enabled     = true
  alert_type  = "universal_ssl_event_type"

  email_integration {
    id = "test@example.com"
  }

  email_integration {
    id = "test2@example.com"
  }
}
