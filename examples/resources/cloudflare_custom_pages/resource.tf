resource "cloudflare_custom_pages" "example_custom_pages" {
  identifier = "023e105f4ecef8ad9ca31a8372d0c353"
  state = "default"
  url = "http://www.example.com"
  zone_id = "zone_id"
}
