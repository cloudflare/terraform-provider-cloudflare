resource "cloudflare_bot_management" "%[1]s" {
  zone_id                = "%[2]s"
  auto_update_model      = true
  enable_js              = false
  suppress_session_score = false
}
