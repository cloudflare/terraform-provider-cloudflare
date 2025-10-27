resource "cloudflare_organization" "%[1]s" {
  name = "%[2]s"

  parent = {
    id = "%[3]s"
  }
}
