# Filter zones using name pattern
data "cloudflare_zones" "%[1]s" {
  account = {
    id = "%[2]s"
  }
  name = "%[3]s"
}