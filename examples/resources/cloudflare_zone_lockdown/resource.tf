# Restrict access to these endpoints to requests from a known IP address range.
resource "cloudflare_zone_lockdown" "example" {
  zone_id     = "d41d8cd98f00b204e9800998ecf8427e"
  paused      = "false"
  description = "Restrict access to these endpoints to requests from a known IP address range"
  urls = [
    "api.mysite.com/some/endpoint*",
  ]
  configurations {
    target = "ip_range"
    value  = "198.51.100.0/16"
  }
}
