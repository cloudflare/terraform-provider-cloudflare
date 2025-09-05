resource "cloudflare_spectrum_application" "%[3]s" {
  zone_id   = "%[1]s"
  protocol  = "tcp/443"
  dns {
    type = "CNAME"
    name = "%[3]s.%[2]s"
  }
  edge_ips {
    type = "dynamic"
    connectivity = "ipv4"
  }
  origin_direct = ["tcp://128.66.0.1:23"]
}