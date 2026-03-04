resource "cloudflare_notification_policy" "%s" {
  account_id  = "%s"
  name        = "%s"
  description = "Test notification policy with traffic anomalies filters"
  enabled     = true
  alert_type  = "traffic_anomalies_alert"

  filters = {
    zones                       = ["%s"]
    selectors                   = ["total"]
    alert_trigger_preferences   = ["zscore_drop"]
    group_by                    = ["zone"]
    where                       = ["(origin_status_code eq 200)"]
  }

  mechanisms = {
    email = [{
      id = "traffic-test@example.com"
    }]
  }
}
