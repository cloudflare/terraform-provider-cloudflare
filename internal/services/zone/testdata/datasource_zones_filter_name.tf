# Filter zones by name
data "cloudflare_zones" "%[1]s" {
  name = "%[2]s"
}