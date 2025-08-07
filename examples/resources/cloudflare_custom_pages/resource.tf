resource "cloudflare_custom_pages" "example_custom_pages" {
  identifier = "ratelimit_block"
  state = "default"
  url = "http://www.example.com"
  zone_id = "zone_id"
}
