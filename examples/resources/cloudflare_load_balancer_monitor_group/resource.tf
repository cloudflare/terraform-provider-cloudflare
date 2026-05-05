resource "cloudflare_load_balancer_monitor_group" "example_load_balancer_monitor_group" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  id = "id"
  description = "Primary datacenter monitors"
  members = [{
    enabled = true
    monitor_id = "monitor_id"
    monitoring_only = false
    must_be_healthy = true
  }]
}
