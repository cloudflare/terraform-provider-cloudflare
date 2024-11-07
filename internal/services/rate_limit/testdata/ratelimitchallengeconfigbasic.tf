
resource "cloudflare_rate_limit" "%[1]s" {
  zone_id = "%[2]s"
  threshold = 1000
  period = 10
  action = [{
    mode = "challenge"
  }]
}