
  resource "cloudflare_healthcheck" "%[2]s" {
    zone_id = "%[1]s"
    description = "Example health check description"
  }