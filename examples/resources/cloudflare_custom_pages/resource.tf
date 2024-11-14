resource "cloudflare_custom_pages" "example" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  type    = "basic_challenge"
  url     = "https://example.com/challenge.html"
  state   = "customized"
}
