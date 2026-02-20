resource "cloudflare_healthcheck" "%s" {
  zone_id = "%s"
  name    = "%s"
  address = "example.com"
  type    = "TCP"

  # v4 flat structure - TCP fields at root level
  method = "connection_established"
  port   = 443

  # Specify check_regions explicitly to avoid API default drift
  check_regions = ["WNAM"]

  interval             = 60
  retries              = 2
  timeout              = 5
  consecutive_fails    = 1
  consecutive_successes = 1
}
