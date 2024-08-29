data "cloudflare_zone" "example" {
  name = "example.com"
}

resource "cloudflare_record" "example" {
  zone_id = data.cloudflare_zone.example.id
  name    = "www"
  value   = "203.0.113.1"
  type    = "A"
  proxied = true
}
