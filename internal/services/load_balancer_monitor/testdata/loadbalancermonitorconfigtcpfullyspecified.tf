
resource "cloudflare_load_balancer_monitor" "test" {
  account_id = "%[1]s"
  type = "tcp"
  method = "connection_established"
  timeout = 9
  interval = 60
  retries = 5
  port = 8080
  description = "this is a very weird tcp load balancer"
}