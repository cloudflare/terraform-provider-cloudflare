# Regionalized hostname record resources are managed independently from the
# Regionalized Hostname resources.
resource "cloudflare_record" "example" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  name    = "example.com"
  value   = "192.0.2.1"
  type    = "A"
  ttl     = 3600
}

# The cloudflare_regional_hostname resource may exist with or without its
# corresponding record resource.
resource "cloudflare_regional_hostname" "example" {
  zone_id    = "0da42c8d2132a9ddaf714f9e7c920711"
  hostname   = "example.com"
  region_key = "eu"
}
