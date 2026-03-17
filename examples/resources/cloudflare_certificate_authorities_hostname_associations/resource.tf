resource "cloudflare_certificate_authorities_hostname_associations" "example_certificate_authorities_hostname_associations" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  hostnames = ["api.example.com"]
  mtls_certificate_id = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
}
