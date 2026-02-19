resource "cloudflare_bot_management" "%s" {
  zone_id           = "%s"
  enable_js         = true
  auto_update_model = true
}
