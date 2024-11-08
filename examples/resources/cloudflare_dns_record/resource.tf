resource "cloudflare_dns_record" "example" {
  zone_id = var.cloudflare_zone_id
  name    = "terraform"
  value   = "192.0.2.1"
  type    = "A"
  ttl     = 3600
}
