resource "cloudflare_managed_transforms" "example_managed_transforms" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  managed_request_headers = [{
    id = "add_cf-bot-score_header"
    enabled = true
  }]
  managed_response_headers = [{
    id = "add_cf-bot-score_header"
    enabled = true
  }]
}
