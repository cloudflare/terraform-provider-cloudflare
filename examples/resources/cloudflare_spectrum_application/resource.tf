resource "cloudflare_spectrum_application" "example" {
  zone_id      = "0da42c8d2132a9ddaf714f9e7c920711"
  protocol     = "tcp/22"
  traffic_type = "direct"

  dns {
    type = "CNAME"
    name = "ssh.example.com"
  }

  origin_direct = [
    "tcp://192.0.2.1:22"
  ]
}
