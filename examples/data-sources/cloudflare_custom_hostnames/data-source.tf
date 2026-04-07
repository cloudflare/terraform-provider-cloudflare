data "cloudflare_custom_hostnames" "example_custom_hostnames" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  id = "0d89c70d-ad9f-4843-b99f-6cc0252067e9"
  certificate_authority = "google"
  custom_origin_server = "origin2.example.com"
  direction = "desc"
  hostname = {
    contain = "example.com"
  }
  hostname_status = "provisioned"
  ssl = 0
  ssl_status = "active"
  wildcard = false
}
