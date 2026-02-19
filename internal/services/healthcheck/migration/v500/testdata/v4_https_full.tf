resource "cloudflare_healthcheck" "%s" {
  zone_id     = "%s"
  name        = "%s"
  address     = "example.com"
  type        = "HTTPS"
  description = "Full HTTPS healthcheck"

  # v4 flat structure - all HTTPS fields at root
  method           = "GET"
  port             = 443
  path             = "/api/health"
  expected_codes   = ["200", "201", "202"]
  expected_body    = "healthy"
  follow_redirects = true
  allow_insecure   = true

  check_regions         = ["WNAM", "ENAM"]
  interval              = 60
  retries               = 2
  timeout               = 10
  consecutive_fails     = 2
  consecutive_successes = 2
  suspended             = false
}
