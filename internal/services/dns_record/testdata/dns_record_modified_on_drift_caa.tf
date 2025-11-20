resource "cloudflare_dns_record" "%[1]s_caa" {
  zone_id = "%[2]s"
  name    = "tf-acctest-caa.%[1]s.%[3]s"
  type    = "CAA"
  ttl     = 3600

  data = {
    flags = "0"
    tag   = "issue"
    value = "letsencrypt.org"
  }
}

# Also test LOC record with data field
resource "cloudflare_dns_record" "%[1]s_loc" {
  zone_id = "%[2]s"
  name    = "tf-acctest-loc.%[1]s.%[3]s"
  type    = "LOC"
  ttl     = 3600

  data = {
    altitude       = 0
    lat_degrees    = 37
    lat_direction  = "N"
    lat_minutes    = 46
    lat_seconds    = 46
    long_degrees   = 122
    long_direction = "W"
    long_minutes   = 23
    long_seconds   = 35
    precision_horz = 0
    precision_vert = 0
    size           = 0
  }
}