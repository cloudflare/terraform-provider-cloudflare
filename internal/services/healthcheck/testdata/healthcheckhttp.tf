
  resource "cloudflare_healthcheck" "%[2]s" {
    zone_id = "%[1]s"
    name = "%[2]s"
    address = "example.com"
    type = "HTTP"
    expected_codes = [
      "200"
    ]
  }