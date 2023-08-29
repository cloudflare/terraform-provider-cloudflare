resource "cloudflare_web_analytics_site" "example" {
  account_id   = "f037e56e89293a057740de681ac9abbe"
  zone_tag     = "0da42c8d2132a9ddaf714f9e7c920711"
  auto_install = true
}

resource "cloudflare_web_analytics_rule" "example" {
  depends_on = [cloudflare_web_analytics_site.example]
  account_id = "f037e56e89293a057740de681ac9abbe"
  ruleset_id = cloudflare_web_analytics_site.example.ruleset_id
  host       = "*"
  paths      = ["/excluded"]
  inclusive  = false
  is_paused  = false
}
