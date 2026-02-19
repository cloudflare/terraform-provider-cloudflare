resource "cloudflare_healthcheck" "%s" {
  zone_id = "%s"
  name    = "%s"
  address = "example.com"
  type    = "HTTP"

  method = "GET"
  port   = 80
  path   = "/health"
  expected_codes = ["200"]

  # Specify check_regions explicitly to avoid API default drift
  check_regions = ["WNAM"]

  # v4 format: header blocks (Set of nested objects)
  header {
    header = "Host"
    values = ["example.com"]
  }

  header {
    header = "User-Agent"
    values = ["Cloudflare-Healthcheck/1.0"]
  }
}
