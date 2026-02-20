resource "cloudflare_healthcheck" "%s" {
  zone_id = "%s"
  name    = "%s"
  address = "example.com"
  type    = "HTTP"

  # Specify check_regions explicitly to avoid API default drift
  check_regions = ["WNAM"]

  # v5 format: nested http_config with header map
  http_config = {
    method         = "GET"
    port           = 80
    path           = "/health"
    expected_codes = ["200"]

    # v5 format: header as map
    header = {
      "Host"       = ["example.com"]
      "User-Agent" = ["Cloudflare-Healthcheck/1.0"]
    }
  }
}
