resource "cloudflare_custom_origin_trust_store" "example_custom_origin_trust_store" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  certificate = <<EOT
  -----BEGIN CERTIFICATE-----
  MIIDdjCCAl6gAwIBAgIJAPnMg0Fs+/B0MA0GCSqGSIb3DQEBCwUAMFsx...
  -----END CERTIFICATE-----

  EOT
}
