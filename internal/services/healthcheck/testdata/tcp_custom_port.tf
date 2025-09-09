resource "cloudflare_healthcheck" "%[2]s" {
  zone_id = "%[1]s"
  name = "%[2]s"
  address = "example.com"
  type = "TCP"
  description = "TCP healthcheck on custom port"
  tcp_config = {
    method = "connection_established"
    port = 8080
  }
  timeout = 15
  retries = 1
}