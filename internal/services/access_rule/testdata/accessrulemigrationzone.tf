resource "cloudflare_access_rule" "%[2]s" {
  zone_id = "%[1]s"
  mode    = "challenge"
  notes   = "Challenge suspicious country"

  configuration {
    target = "country"
    value  = "TV"
  }
}
