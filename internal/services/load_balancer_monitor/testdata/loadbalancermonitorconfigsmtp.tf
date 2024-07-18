
resource "cloudflare_load_balancer_monitor" "%[1]s" {
  account_id = "%[2]s"
  type = "smtp"
  timeout = 2
  interval = 60
  retries = 5
  port = 8080
  description = "test setup smtp"
}