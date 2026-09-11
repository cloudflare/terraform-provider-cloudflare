resource "cloudflare_email_sending_subdomain" "example_email_sending_subdomain" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  name = "sub.example.com"
}
