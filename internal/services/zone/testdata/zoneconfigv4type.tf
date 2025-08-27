resource "cloudflare_zone" "%[1]s" {
  zone       = "%[2]s"
  account_id = "%[3]s"
  type       = "%[4]s"
  plan       = "enterprise"
}