resource "cloudflare_notification_policy" "%s" {
  account_id = "%s"
  name       = "%s"
  enabled    = true
  alert_type = "universal_ssl_event_type"

  mechanisms = {
    email = [{
      id = "test-minimal@example.com"
    }]
  }
}
