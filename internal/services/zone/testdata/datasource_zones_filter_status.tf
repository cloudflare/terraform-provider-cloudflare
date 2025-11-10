# Filter zones by account and status
data "cloudflare_zones" "%[1]s" {
  account = {
    id = "%[2]s"
  }
  status = "%[3]s"
}