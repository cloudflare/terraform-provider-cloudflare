resource "cloudflare_custom_pages" "%[1]s" {
  zone_id = "%[2]s"
  type    = "%[3]s"
  state   = "%[4]s"
  url     = "%[5]s"
}