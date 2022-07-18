resource "cloudflare_custom_hostname" "example" {
  zone_id  = "1d5fdc9e88c8a8c4518b068cd94331fe"
  hostname = "hostname.example.com"
  ssl {
    method = "txt"
  }
}
