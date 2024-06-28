# HTTP
resource "cloudflare_load_balancer_monitor" "example" {
  account_id     = "f037e56e89293a057740de681ac9abbe"
  type           = "http"
  expected_body  = "alive"
  expected_codes = "2xx"
  method         = "GET"
  timeout        = 7
  path           = "/health"
  interval       = 60
  retries        = 5
  description    = "example http load balancer"
  header {
    header = "Host"
    values = ["example.com"]
  }
  allow_insecure   = false
  follow_redirects = true
  probe_zone       = "example.com"
}

# TCP Monitor
resource "cloudflare_load_balancer_monitor" "example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  type        = "tcp"
  method      = "connection_established"
  timeout     = 7
  port        = 8080
  interval    = 60
  retries     = 5
  description = "example tcp load balancer"
}
