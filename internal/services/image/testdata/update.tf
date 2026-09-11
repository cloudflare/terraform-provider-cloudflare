resource "cloudflare_image" "%[1]s" {
  account_id = "%[2]s"
  id         = "%[1]s"
  url        = "%[3]s"
  metadata = jsonencode({
    source = "terraform"
    env    = "prod"
  })
}
