# Reginoalized hostname record resources are
# managed independently from the Reginoalized Hostname resources
resource "cloudflare_record" "example" {
  zone_id = var.cloudflare_zone_id
  name    = "example.com"
  value   = "192.0.2.1"
  type    = "A"
  ttl     = 3600
}

# The regional_hostname resource may exist with or without its
# corresponding record resource.
resource cloudflare_regional_hostname "regional_hostname_example" {
  zone_id = var.cloudflare_zone_id
  hostname   = "example.com"
  region_key = "eu"
}
