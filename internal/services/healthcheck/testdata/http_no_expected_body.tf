resource "cloudflare_healthcheck" "%[2]s" {
  zone_id = "%[1]s"
  name = "%[2]s"
  address = "example.com"
  type = "HTTP"
  check_regions = ["WNAM"]
  http_config = {
    method = "GET"
    path = "/"
    port = 80
    expected_codes = ["200"]
  }
}
