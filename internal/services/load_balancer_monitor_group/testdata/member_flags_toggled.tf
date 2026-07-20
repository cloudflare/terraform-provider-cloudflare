resource "cloudflare_load_balancer_monitor" "%[1]s_a" {
  account_id     = "%[2]s"
  type           = "https"
  expected_codes = "2xx"
  expected_body  = "alive"
}

resource "cloudflare_load_balancer_monitor_group" "%[1]s" {
  account_id  = "%[2]s"
  description = "tf-acc basic %[1]s"
  members = [
    {
      enabled         = false
      monitor_id      = cloudflare_load_balancer_monitor.%[1]s_a.id
      monitoring_only = true
      must_be_healthy = false
    },
  ]
}
