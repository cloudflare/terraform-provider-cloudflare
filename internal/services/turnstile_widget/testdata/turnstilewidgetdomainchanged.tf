
resource "cloudflare_turnstile_widget" "%[1]s" {
  account_id     = "%[2]s"
  name           = "%[1]s"
  bot_fight_mode = false
  domains        = ["b.example.com", "a.example.com", "d.example.com"]
  mode           = "invisible"
  region         = "world"
}
