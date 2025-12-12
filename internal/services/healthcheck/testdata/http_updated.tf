resource "cloudflare_healthcheck" "%[2]s" {
  zone_id               = "%[1]s"
  name                  = "%[2]s-updated"
  address               = "example.com"
  type                  = "HTTP"
  description           = "HTTP healthcheck updated"
  check_regions         = ["WNAM", "ENAM"]
  consecutive_fails     = 3
  consecutive_successes = 3
  interval              = 120
  retries               = 3
  suspended             = true
  timeout               = 10
  http_config = {
    allow_insecure   = true
    expected_body    = "HEALTHY"
    expected_codes   = ["200", "201"]
    follow_redirects = true
    header = {
      "Host"            = ["example.com"]
      "X-Custom-Header" = ["value1"]
    }
    method = "HEAD"
    path   = "/ping"
    port   = 8080
  }
}
