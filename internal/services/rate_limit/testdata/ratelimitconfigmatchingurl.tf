
resource "cloudflare_rate_limit" "%[1]s" {
  zone_id = "%[2]s"
  threshold = 1000
  period = 10
  match = [{
    request = [{
      url_pattern = "%[3]s/tfacc-url-%[1]s"
    }]
  }]
  action = [{
    mode = "simulate"
    timeout = 86400
  }]
}