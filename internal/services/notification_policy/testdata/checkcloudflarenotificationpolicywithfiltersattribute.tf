
  resource "cloudflare_notification_policy" "%[1]s" {
    name        = "workers usage notification"
    account_id  = "%[2]s"
    description = "test description"
    enabled     =  true
    alert_type  = "billing_usage_alert"
    email_integration = [{
      name =  ""
      id   =  "test@example.com"
    }]
    filters = {
  product = [
        "worker_requests",
      ]
	  limit = ["100"]
}
  }