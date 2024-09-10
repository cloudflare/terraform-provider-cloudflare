resource "cloudflare_access_custom_page" "example" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  name        = "example"
  type        = "forbidden"
  custom_html = "<html><body><h1>Forbidden</h1></body></html>"
}
