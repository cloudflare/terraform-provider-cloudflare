resource "cloudflare_turnstile_widget" "%[1]s" {
  account_id     = "%[2]s"
  name           = "%[1]s-updated"
  bot_fight_mode = false
  domains        = ["example.com", "test.example.com"]
  mode           = "invisible"
  region         = "world"
}
