
  resource "cloudflare_notification_policy" "%[1]s" {
    name        = "traffic anomalies alert"
    account_id  = "%[2]s"
    description = "test description"
    enabled     =  true
    alert_type  = "incident_alert"
	email_integration = [{
      name =  ""
      id   =  "test@example.com"
    }]
    filters = {
  affected_components = ["API"]
}
  }