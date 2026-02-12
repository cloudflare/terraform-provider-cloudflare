resource "cloudflare_spectrum_application" "%s" {
  zone_id   = "%s"
  protocol  = "tcp/3306"
  dns {
    type = "CNAME"
    name = "%s.%s"
  }
  origin_direct = ["tcp://128.66.0.2:3306"]
  origin_port = 3306
}
