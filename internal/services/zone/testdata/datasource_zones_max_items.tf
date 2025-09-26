# Zones with max_items limit
data "cloudflare_zones" "%[1]s" {
  account = {
    id = "%[2]s"
  }
  max_items = %[3]d
}