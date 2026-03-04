resource "cloudflare_turnstile_widget" "%s" {
  account_id     = "%s"
  name           = "%s"
  domains        = ["example.com", "test.example.org"]
  mode           = "invisible"
  region         = "world"
  bot_fight_mode = false
  offlabel       = false
}
