resource "cloudflare_custom_pages" "%[1]s" {
  account_id = "%[2]s"
  identifier = "%[3]s"
  state      = "%[4]s"
  url        = "%[5]s"
}