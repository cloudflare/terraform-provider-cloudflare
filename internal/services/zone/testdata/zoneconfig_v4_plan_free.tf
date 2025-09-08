resource "cloudflare_zone" "%[1]s" {
  zone       = "%[2]s"
  account_id = "%[3]s"
  plan       = "free"
  jump_start = false
}