resource "cloudflare_r2_custom_domain" "example_r2_custom_domain" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  bucket_name = "example-bucket"
  domain = "domain"
  enabled = true
  zone_id = "zoneId"
  min_tls = "1.0"
}
