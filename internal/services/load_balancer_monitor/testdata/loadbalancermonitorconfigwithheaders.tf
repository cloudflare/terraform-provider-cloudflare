
resource "cloudflare_load_balancer_monitor" "%[1]s" {
  account_id = "%[3]s"
  expected_body = "dead"
  expected_codes = "5xx"
  method = "HEAD"
  timeout = 9
  path = "/custom"
  interval = 60
  retries = 5
  port = 8080
  description = "this is a very weird load balancer"
  header ={
    header = "Host"
    values = ["%[2]s"]
  }
}
