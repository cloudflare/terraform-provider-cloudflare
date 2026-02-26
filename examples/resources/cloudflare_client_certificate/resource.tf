resource "cloudflare_client_certificate" "example_client_certificate" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  csr = <<EOT
  -----BEGIN CERTIFICATE REQUEST-----
  MIICY....
  -----END CERTIFICATE REQUEST-----
  EOT
  validity_days = 3650
}
