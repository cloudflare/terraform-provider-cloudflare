resource "cloudflare_user_agent_blocking_rule" "example_1" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  mode        = "js_challenge"
  paused      = false
  description = "My description 1"
  configuration {
    target = "ua"
    value  = "Chrome"
  }
}

resource "cloudflare_user_agent_blocking_rule" "example_2" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  mode        = "challenge"
  paused      = true
  description = "My description 22"
  configuration {
    target = "ua"
    value  = "Mozilla"
  }
}
