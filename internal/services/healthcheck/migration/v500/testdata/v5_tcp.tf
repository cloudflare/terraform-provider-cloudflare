resource "cloudflare_healthcheck" "%s" {
  zone_id = "%s"
  name    = "%s"
  address = "example.com"
  type    = "TCP"

  # v5 nested structure - TCP fields in tcp_config
  tcp_config = {
    method = "connection_established"
    port   = 443
  }

  # Specify check_regions explicitly to avoid API default drift
  check_regions = ["WNAM"]

  interval              = 60
  retries               = 2
  timeout               = 5
  consecutive_fails     = 1
  consecutive_successes = 1
}
