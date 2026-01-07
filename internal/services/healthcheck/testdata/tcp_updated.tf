resource "cloudflare_healthcheck" "%[3]s" {
  zone_id               = "%[1]s"
  name                  = "%[2]s"
  address               = "example.com"
  type                  = "TCP"
  description           = "TCP healthcheck updated"
  check_regions         = ["WNAM", "ENAM"]
  consecutive_fails     = 3
  consecutive_successes = 3
  interval              = 120
  retries               = 3
  suspended             = true
  timeout               = 15
  tcp_config = {
    method = "connection_established"
    port   = 8080
  }
}
