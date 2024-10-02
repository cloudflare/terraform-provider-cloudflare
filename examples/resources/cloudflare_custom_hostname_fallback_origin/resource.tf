resource "cloudflare_custom_hostname_fallback_origin" "example" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  origin  = "fallback.example.com"
}
