# Look up zone by name only
data "cloudflare_zone" "%[1]s" {
  filter = {
    name = "%[2]s"
  }
}