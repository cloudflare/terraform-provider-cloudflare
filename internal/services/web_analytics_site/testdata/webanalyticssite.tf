
resource "cloudflare_web_analytics_site" "%[1]s" {
  account_id    = "%[2]s"
  host          = "%[3]s"
  auto_install  = false
}
