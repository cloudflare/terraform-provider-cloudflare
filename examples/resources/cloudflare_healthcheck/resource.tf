# HTTPS Healthcheck
resource "cloudflare_healthcheck" "http_health_check" {
  zone_id = var.cloudflare_zone_id
  name = "http-health-check"
  description = "example http health check"
  address = "example.com"
  suspended = false
  check_regions = [
    "WEU",
    "EEU"
  ]
  type = "HTTPS"
  port = 443
  method = "GET"
  path = "/health"
  expected_body = "alive"
  expected_codes = [
    "2xx",
    "301"
  ]
  follow_redirects = true
  allow_insecure = false
  header {
    header = "Host"
    values = ["example.com"]
  }
  timeout = 10
  retries = 2
  interval = 60
  consecutive_fails = 3
  consecutive_successes = 2
}

# TCP Healthcheck
resource "cloudflare_healthcheck" "tcp_health_check" {
  zone_id = var.cloudflare_zone_id
  name = "tcp-health-check"
  description = "example tcp health check"
  address = "example.com"
  suspended = false
  check_regions = [
    "WEU",
    "EEU"
  ]
  type = "TCP"
  port = 22
  method = "connection_established"
  timeout = 10
  retries = 2
  interval = 60
  consecutive_fails = 3
  consecutive_successes = 2
}
