data "cloudflare_authenticated_origin_pulls" "example_authenticated_origin_pulls" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  hostname = "app.example.com"
}
