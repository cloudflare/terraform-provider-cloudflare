# Basic zones lookup with account filter
data "cloudflare_zones" "%[1]s" {
  account = {
    id = "%[2]s"
  }
}