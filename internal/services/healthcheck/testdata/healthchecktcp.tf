
  resource "cloudflare_healthcheck" "%[3]s" {
    zone_id = "%[1]s"
    name = "%[2]s"
    address = "example.com"
    type = "TCP"
    method = "connection_established"
    port = 80
  }