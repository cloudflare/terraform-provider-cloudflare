
resource "cloudflare_notification_policy" "%[1]s" {
  name        = "%[2]s"
  account_id  = "%[4]s"
  description = "%[3]s"
  enabled     =  true
  alert_type  = "billing_usage_alert"
  mechanisms = {
    "email": [{"id": "test@example.com"}]
  }
  filters = {
    product = [
      "worker_requests",
    ]
    limit = ["100"]
  }
}