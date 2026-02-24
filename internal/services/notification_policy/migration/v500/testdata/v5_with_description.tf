resource "cloudflare_notification_policy" "%s" {
  account_id  = "%s"
  name        = "%s"
  description = "This is a comprehensive test description field for migration testing purposes"
  enabled     = true
  alert_type  = "universal_ssl_event_type"

  mechanisms = {
    email = [{
      id = "description-test@example.com"
    }]
  }
}
