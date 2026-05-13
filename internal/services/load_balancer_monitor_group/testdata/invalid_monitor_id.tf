resource "cloudflare_load_balancer_monitor_group" "%[1]s" {
  account_id  = "%[2]s"
  description = "tf-acc invalid-monitor %[1]s"
  members = [
    {
      enabled         = true
      monitor_id      = "00000000000000000000000000000000"
      monitoring_only = false
      must_be_healthy = true
    },
  ]
}
