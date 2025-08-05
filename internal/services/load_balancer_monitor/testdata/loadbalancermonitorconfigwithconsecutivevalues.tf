resource "cloudflare_load_balancer_monitor" "%s" {
  account_id = "%s"
  type       = "http"
  method     = "GET"
  path       = "/"
  expected_codes = "200"
  consecutive_up = 0
  consecutive_down = 0
  port = 0
}
