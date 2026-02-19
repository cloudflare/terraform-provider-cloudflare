resource "cloudflare_spectrum_application" "%s" {
  zone_id   = "%s"
  protocol  = "tcp/443"
  dns {
    type = "CNAME"
    name = "%s.%s"
  }
  edge_ips {
    type = "dynamic"
    connectivity = "ipv4"
  }
  origin_direct = ["tcp://128.66.0.1:23"]
}
