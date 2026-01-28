resource "cloudflare_account" "%[1]s" {
  name = "%[2]s"
  settings = {
      enforce_twofactor = "%[3]t"
  }
}