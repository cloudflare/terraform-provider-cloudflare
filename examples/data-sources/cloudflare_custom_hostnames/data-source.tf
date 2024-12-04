data "cloudflare_custom_hostnames" "example_custom_hostnames" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  id = "0d89c70d-ad9f-4843-b99f-6cc0252067e9"
  direction = "asc"
  hostname = "app.example.com"
  ssl = 0
}
