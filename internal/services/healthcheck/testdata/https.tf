resource "cloudflare_healthcheck" "%[2]s" {
  zone_id               = "%[1]s"
  name                  = "%[2]s"
  address               = "example.com"
  type                  = "HTTPS"
  description           = "HTTPS healthcheck"
  check_regions         = ["WNAM", "ENAM"]
  consecutive_fails     = 2
  consecutive_successes = 2
  interval              = 60
  retries               = 2
  timeout               = 10
  suspended             = false
  http_config = {
    allow_insecure   = false
    expected_body    = "OK"
    expected_codes   = ["200", "201"]
    follow_redirects = true
    header = {
      "Host" = ["example.com"]
    }
    method = "GET"
    path   = "/health"
    port   = 443
  }
}
