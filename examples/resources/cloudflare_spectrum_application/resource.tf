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

  edge_ips {
    type = "static"
    ips  = ["203.0.113.1", "203.0.113.2"]
  }
}
