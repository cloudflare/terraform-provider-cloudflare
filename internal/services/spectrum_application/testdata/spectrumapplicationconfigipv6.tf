resource "cloudflare_spectrum_application" "%[3]s" {
  zone_id  = "%[1]s"
  protocol = "tcp/22"

  dns = {
    type = "CNAME"
    name = "%[3]s.%[2]s"
  }

  origin_direct = ["tcp://128.66.0.10:22"]
  origin_port   = 22
  
  # Test IPv6 edge IP connectivity
  edge_ips = {
    type = "dynamic"
    connectivity = "ipv6"
  }
}
