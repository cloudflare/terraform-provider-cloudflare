
resource "cloudflare_load_balancer_monitor" "%[3]s" {
  account_id = "%[2]s"
  expected_body = "dead"
  expected_codes = "5xx"
  method = "HEAD"
  timeout = 9
  path = "/custom"
  interval = 60
  retries = 5
  consecutive_up = 2
  consecutive_down = 2
  port = 8080
  description = "this is a very weird load balancer"
  probe_zone = "%[1]s"
  header = {
    Header = ["Host"]
    Values = ["%[1]s"]
  }
}
