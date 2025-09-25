resource "cloudflare_spectrum_application" "%[3]s" {
  zone_id   = "%[1]s"
  protocol  = "tcp/3306"
  dns {
    type = "CNAME"
    name = "%[3]s.%[2]s"
  }
  origin_port_range {
    start = 3306
    end   = 3310
  }
  origin_direct = ["tcp://128.66.0.1:23"]
}