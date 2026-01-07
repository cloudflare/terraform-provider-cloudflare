resource "cloudflare_healthcheck" "%[3]s" {
  zone_id               = "%[1]s"
  name                  = "%[2]s"
  address               = "example.com"
  type                  = "TCP"
  description           = "TCP healthcheck"
  check_regions         = ["WNAM"]
  consecutive_fails     = 2
  consecutive_successes = 2
  interval              = 60
  retries               = 2
  suspended             = false
  timeout               = 10
  tcp_config = {
    method = "connection_established"
    port   = 80
  }
}
