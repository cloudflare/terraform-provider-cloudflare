resource "cloudflare_healthcheck" "%s" {
  zone_id = "%s"
  name    = "%s"
  address = "example.com"
  type    = "HTTP"

  # Nested structure in v5 - HTTP fields in http_config
  http_config = {
    method           = "GET"
    port             = 80
    path             = "/health"
    expected_codes   = ["200", "201"]
    expected_body    = "OK"
    follow_redirects = false
    allow_insecure   = false
  }

  # Specify check_regions explicitly to avoid API default drift
  check_regions = ["WNAM"]

  interval              = 60
  retries               = 2
  timeout               = 5
  consecutive_fails     = 1
  consecutive_successes = 1
}
