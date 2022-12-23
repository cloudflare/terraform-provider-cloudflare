# Runs the specified worker script for all URLs that match `example.com/*`
resource "cloudflare_worker_route" "my_route" {
  zone_id     = "d41d8cd98f00b204e9800998ecf8427e"
  pattern     = "example.com/*"
  script_name = cloudflare_worker_script.my_script.name
}

resource "cloudflare_worker_script" "my_script" {
  # see "cloudflare_worker_script" documentation ...
}
