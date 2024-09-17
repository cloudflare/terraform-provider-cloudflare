
  resource "cloudflare_healthcheck" "%[2]s" {
    zone_id = "%[1]s"
    name = "%[2]s"
    address = "example.com"
    type = "HTTP"
    http_config = {
      expected_codes = [
        "200"
      ]
    }
  }
