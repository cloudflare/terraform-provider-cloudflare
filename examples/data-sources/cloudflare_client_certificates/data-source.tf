data "cloudflare_client_certificates" "example_client_certificates" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  limit = 10
  offset = 10
  status = "all"
}
