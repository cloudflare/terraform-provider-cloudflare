
resource "cloudflare_access_rule" "%[6]s" {
  zone_id = "%[1]s"
  notes = "%[3]s"
  mode = "%[2]s"
  configuration = {
  target = "%[4]s"
    value = "%[5]s"
}
}