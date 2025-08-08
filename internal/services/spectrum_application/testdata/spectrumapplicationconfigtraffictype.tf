resource "cloudflare_spectrum_application" "%[3]s" {
  zone_id  = "%[1]s"
  protocol = "tcp/80"

  dns = {
    type = "CNAME"
    name = "%[3]s.%[2]s"
  }

  origin_direct = ["tcp://128.66.0.9:80"]
  origin_port   = 80
  
  # Test traffic type configuration
  traffic_type = "%[4]s"  # direct, http, https
  
  edge_ips = {
    type = "dynamic"
    connectivity = "all"
  }
}
