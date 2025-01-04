
resource "cloudflare_user_agent_blocking_rule" "%[1]s" {
  zone_id     = "%[2]s"
  mode        = "%[3]s"
  paused      = false
  description = "My description"
  configuration = {
    target = "ua"
    value  = "Mozilla"
  }
}
