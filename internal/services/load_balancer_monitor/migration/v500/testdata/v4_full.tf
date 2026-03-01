resource "cloudflare_load_balancer_monitor" "%s" {
  account_id        = "%s"
  type              = "https"
  method            = "GET"
  path              = "/api/health"
  port              = 8443
  interval          = 60
  retries           = 3
  timeout           = 10
  allow_insecure    = true
  follow_redirects  = true
  expected_codes    = "2xx"
  expected_body     = "OK"
  description       = "Production API health check"
  consecutive_up    = 2
  consecutive_down  = 3
  probe_zone        = "%s"

  header {
    header = "Host"
    values = ["%s"]
  }
}
