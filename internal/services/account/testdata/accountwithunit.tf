resource "cloudflare_account" "%[1]s" {
  name = "%[2]s"
  type = "enterprise"
  unit = {
    id   = "%[3]s"
  }
}