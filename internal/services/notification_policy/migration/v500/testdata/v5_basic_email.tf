resource "cloudflare_notification_policy" "%s" {
  account_id  = "%s"
  name        = "%s"
  description = "Test notification policy with email integration"
  enabled     = true
  alert_type  = "universal_ssl_event_type"

  mechanisms = {
    email = [{
      id = "test@example.com"
    }]
  }
}
