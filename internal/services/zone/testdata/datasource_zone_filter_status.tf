# Look up zone with filter by name and status
data "cloudflare_zone" "%[1]s" {
  filter = {
    name   = "%[2]s"
    status = "%[3]s"
  }
}