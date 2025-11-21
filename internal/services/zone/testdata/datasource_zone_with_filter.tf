# Look up zone with filter by name and account
data "cloudflare_zone" "%[1]s" {
  filter = {
    name = "%[2]s"
    account = {
      id = "%[3]s"
    }
  }
}