
  resource "cloudflare_notification_policy" "%[1]s" {
    name        = "%[2]s"
    account_id  = "%[4]s"
    description = "%[3]s"
    enabled     =  true
    alert_type  = "universal_ssl_event_type"
    email_integration = [{
      name =  ""
      id   =  "test@example.com"
    },
    {
    name =  ""
      id   =  "test2@example.com"
    }]
  }