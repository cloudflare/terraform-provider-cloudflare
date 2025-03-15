resource "cloudflare_page_rule" "example_page_rule" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  target = "example.com/*"
  priority = 1
  status = "active"
  actions = {
    forwarding_url = "https://example.com/foo"
    status_code = 301
  }
}
