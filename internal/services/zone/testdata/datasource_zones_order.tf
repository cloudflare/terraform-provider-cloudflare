# Zones with ordering and direction
data "cloudflare_zones" "%[1]s" {
  account = {
    id = "%[2]s"
  }
  order     = "%[3]s"
  direction = "%[4]s"
}