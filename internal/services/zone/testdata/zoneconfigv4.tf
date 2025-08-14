resource "cloudflare_zone" "%[1]s" {
  zone       = "%[2]s"
  account_id = "%[3]s"
  paused     = false
  plan       = "free"
  type       = "full"
  jump_start = true
}