
resource "cloudflare_web_analytics_site" "%[1]s" {
  account_id   = "%[2]s"
  zone_tag     = "%[3]s"
  auto_install = true
}

resource "cloudflare_web_analytics_rule" "%[1]s" {
  depends_on = [cloudflare_web_analytics_site.%[1]s]
  account_id = "%[2]s"
  ruleset_id = cloudflare_web_analytics_site.%[1]s.ruleset.id
  host       = "%[3]s"
  paths      = ["/excluded"]
  inclusive  = false
  is_paused  = false
}
