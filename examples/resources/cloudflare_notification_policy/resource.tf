resource "cloudflare_notification_policy" "example" {
  account_id  = "c4a7362d577a6c3019a474fd6f485821"
  name        = "Policy for SSL notification events"
  description = "Notification policy to alert when my SSL certificates are modified"
  enabled     = true
  alert_type  = "universal_ssl_event_type"

  email_integration {
    id = "myemail@example.com"
  }

  webhooks_integration {
    id = "1860572c5d964d27aa0f379d13645940"
  }

  pagerduty_integration {
    id = "850129d136459401860572c5d964d27k"
  }
}

### With Filters
resource "cloudflare_notification_policy" "example" {
  account_id  = "c4a7362d577a6c3019a474fd6f485821"
  name        = "Policy for Healthcheck notification"
  description = "Notification policy to alert on unhealthy Healthcheck status"
  enabled     = true
  alert_type  = "health_check_status_notification"

  email_integration {
    id = "myemail@example.com"
  }

  webhooks_integration {
    id = "1860572c5d964d27aa0f379d13645940"
  }

  pagerduty_integration {
    id = "850129d136459401860572c5d964d27k"
  }

  filters {
    health_check_id = ["699d98642c564d2e855e9661899b7252"]
    status          = ["Unhealthy"]
  }
}
