resource "cloudflare_zone" "example" {
  zone = "example.com"
}

resource "cloudflare_zone_dnssec" "example" {
  zone_id = cloudflare_zone.example.id
}
