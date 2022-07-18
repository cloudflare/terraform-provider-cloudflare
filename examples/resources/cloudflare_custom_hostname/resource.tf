resource "cloudflare_custom_hostname" "example" {
  zone_id  = "0da42c8d2132a9ddaf714f9e7c920711"
  hostname = "hostname.example.com"
  ssl {
    method = "txt"
  }
}
