# Restrict access to these endpoints to requests from a known IP address range.
resource "cloudflare_zone_lockdown" "example" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  paused      = "false"
  description = "Restrict access to these endpoints to requests from a known IP address range"
  urls = [
    "api.mysite.com/some/endpoint*",
  ]
  configurations {
    target = "ip_range"
    value  = "192.0.2.0/24"
  }
}
