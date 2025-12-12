resource "cloudflare_healthcheck" "%[2]s" {
  zone_id               = "%[1]s"
  name                  = "%[2]s-updated"
  address               = "example.com"
  type                  = "HTTPS"
  description           = "HTTPS healthcheck updated"
  check_regions         = ["WNAM"]
  consecutive_fails     = 3
  consecutive_successes = 3
  interval              = 120
  retries               = 3
  timeout               = 5
  suspended             = true
  http_config = {
    allow_insecure   = true
    expected_body    = "HEALTHY"
    expected_codes   = ["200"]
    follow_redirects = false
    header = {
      "Host"            = ["example.com"]
      "X-Custom-Header" = ["value1"]
    }
    method = "HEAD"
    path   = "/status"
    port   = 8443
  }
}
