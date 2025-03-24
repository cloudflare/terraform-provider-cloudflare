
resource "cloudflare_notification_policy" "%[1]s" {
  name        = "test SSL policy from terraform provider"
  account_id  = "%[2]s"
  description = "test description"
  enabled     =  true
  alert_type  = "universal_ssl_event_type"
  mechanisms = {
    "email": [{"id": "test@example.com"},{"id": "test2@example.com"}]
  }
}