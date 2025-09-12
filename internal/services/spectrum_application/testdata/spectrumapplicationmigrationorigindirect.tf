resource "cloudflare_spectrum_application" "%[3]s" {
  zone_id   = "%[1]s"
  protocol  = "tcp/3306"
  dns {
    type = "CNAME"
    name = "%[3]s.%[2]s"
  }
  origin_direct = ["tcp://128.66.0.2:3306"]
  origin_port = 3306
}