resource "cloudflare_spectrum_application" "%[3]s" {
  zone_id  = "%[1]s"
  protocol = "tcp/22"

  dns = {
    type = "CNAME"
    name = "%[3]s.%[2]s"
  }

  origin_direct = ["tcp://128.66.0.8:22"]
  origin_port   = 22
  
  # Test Argo Smart Routing configuration
  argo_smart_routing = %[4]s  # true or false
  traffic_type = "direct"     # Required for Argo Smart Routing
  
  edge_ips = {
    type = "dynamic"
    connectivity = "all"
  }
}
