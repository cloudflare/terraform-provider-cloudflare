resource "cloudflare_workers_route" "example_workers_route" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  id = "023e105f4ecef8ad9ca31a8372d0c353"
  pattern = "example.com/*"
  script = "my-workers-script"
}
