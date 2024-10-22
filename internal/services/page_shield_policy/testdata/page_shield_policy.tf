resource "cloudflare_page_shield_policy" "%[1]s" {
  zone_id = "%[2]s"
  description = "%[3]s"
  action = "%[4]s"
  expression = "%[5]s"
  enabled = "%[6]s"
  value = "%[7]s"
}
