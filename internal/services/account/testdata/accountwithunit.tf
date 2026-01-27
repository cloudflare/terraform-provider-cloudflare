resource "cloudflare_account" "%[1]s" {
  name = "%[2]s"
  unit = {
    id   = "%[3]s"
  }
}