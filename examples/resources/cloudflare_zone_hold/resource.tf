resource "cloudflare_zone_hold" "example" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  hold    = true
}
