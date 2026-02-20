resource "cloudflare_bot_management" "%s" {
  zone_id                = "%s"
  enable_js              = true
  auto_update_model      = true
  suppress_session_score = false
  ai_bots_protection     = "block"
}
