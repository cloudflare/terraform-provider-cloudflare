
resource "cloudflare_load_balancer_monitor" "%[1]s" {
  account_id = "%[2]s"
  type = "icmp_ping"
  timeout = 2
  interval = 60
  retries = 5
  description = "test setup icmp_ping"
}