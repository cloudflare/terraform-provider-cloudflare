resource "cloudflare_bot_management" "example_bot_management" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  ai_bots_protection = "block"
  cf_robots_variant = "policy_only"
  crawler_protection = "enabled"
  enable_js = true
  fight_mode = true
  is_robots_txt_managed = false
}
