resource "cloudflare_workers_route" "example_workers_route" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  pattern = "example.net/*"
  script = "this-is_my_script-01"
}
