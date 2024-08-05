# Runs the specified worker script for all URLs that match `example.com/*`
resource "cloudflare_workers_route" "my_route" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  pattern     = "example.com/*"
  script_name = cloudflare_worker_script.my_script.name
}

resource "cloudflare_workers_script" "my_script" {
  # see "cloudflare_workers_script" documentation ...
}
