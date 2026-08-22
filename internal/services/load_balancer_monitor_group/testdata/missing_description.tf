resource "cloudflare_load_balancer_monitor" "%[1]s_a" {
  account_id     = "%[2]s"
  type           = "https"
  expected_codes = "2xx"
  expected_body  = "alive"
}

resource "cloudflare_load_balancer_monitor_group" "%[1]s" {
  account_id = "%[2]s"
  members = [
    {
      enabled         = true
      monitor_id      = cloudflare_load_balancer_monitor.%[1]s_a.id
      monitoring_only = false
      must_be_healthy = true
    },
  ]
}
