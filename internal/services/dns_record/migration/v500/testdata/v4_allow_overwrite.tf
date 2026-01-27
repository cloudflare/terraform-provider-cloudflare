resource "cloudflare_record" "%[1]s" {
  zone_id         = "%[2]s"
  name            = "%[3]s"
  type            = "A"
  value           = "192.0.2.2"
  ttl             = 3600
  allow_overwrite = true
}
