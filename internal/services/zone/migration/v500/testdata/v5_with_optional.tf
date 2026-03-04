resource "cloudflare_zone" "%[1]s" {
  name   = "%[2]s"
  paused = false
  type   = "full"
  account = {
    id = "%[3]s"
  }
}
