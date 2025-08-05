resource "cloudflare_load_balancer_monitor" "%s" {
  account_id = "%s"
  type       = "http"
  method     = "GET"
  path       = "/"
  expected_codes = "200"
  description = "%s"
  # Non-zero values for consecutive_up, consecutive_down, and port
  consecutive_up = 2
  consecutive_down = 2
  port = 8080
}
