resource "cloudflare_healthcheck" "%[2]s" {
  zone_id               = "%[1]s"
  name                  = "%[2]s"
  address               = "example.com"
  type                  = "HTTP"
  description           = "HTTP healthcheck"
  check_regions         = ["WNAM"]
  consecutive_fails     = 2
  consecutive_successes = 2
  interval              = 60
  retries               = 2
  suspended             = false
  timeout               = 5
  http_config = {
    allow_insecure   = false
    expected_body    = "OK"
    expected_codes   = ["200"]
    follow_redirects = false
    header = {
      "Host" = ["example.com"]
    }
    method = "GET"
    path   = "/health"
    port   = 80
  }
}
