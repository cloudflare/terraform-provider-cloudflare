resource "cloudflare_spectrum_application" "%[3]s" {
  zone_id  = "%[1]s"
  protocol = "udp/53"

  dns = {
    type = "CNAME"
    name = "%[3]s.%[2]s"
  }

  origin_direct = ["udp://128.66.0.11:53"]
  origin_port   = 53
  
  # Test UDP protocol configuration
  traffic_type = "direct"
  
  edge_ips = {
    type = "dynamic"
    connectivity = "all"
  }
}
