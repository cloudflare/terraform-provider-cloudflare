resource "cloudflare_custom_page_asset" "example_custom_page_asset" {
  description = "Custom 500 error page"
  name = "my_custom_error_page"
  url = "https://example.com/error.html"
  zone_id = "zone_id"
}
