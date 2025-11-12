
resource "cloudflare_dns_record" "%[3]s" {
  zone_id = "%[1]s"
  name = "%[2]s"
  data = {
  lat_degrees    = "37"
    lat_minutes    = "46"
    lat_seconds    = 46.000
    lat_direction  = "N"
    long_degrees   = "122"
    long_minutes   = "23"
    long_seconds   = 35.000
    long_direction = "W"
    altitude       = 0.00
    size           = 100.00
    precision_horz = 0.00
    precision_vert = 0.00
}
  type = "LOC"
  ttl = 3600
}
