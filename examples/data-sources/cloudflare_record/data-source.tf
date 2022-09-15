data "cloudflare_record" "example" {
  zone_id = var.zone_id
  hostname = "example.com"
}