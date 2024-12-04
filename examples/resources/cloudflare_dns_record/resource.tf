resource "cloudflare_dns_record" "example_dns_record" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  content = "198.51.100.4"
  type = "A"
}
