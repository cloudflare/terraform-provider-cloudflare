resource "cloudflare_spectrum_application" "%[3]s" {
  zone_id  = "%[1]s"
  protocol = "tcp/22"

  dns = {
    type = "CNAME"
    name = "%[3]s.%[2]s"
  }

  origin_direct = ["tcp://128.66.0.7:22"]
  origin_port   = 22
  
  # Test IP firewall configuration
  ip_firewall = %[4]s  # true or false
  
  edge_ips = {
    type = "dynamic"
    connectivity = "all"
  }
}
