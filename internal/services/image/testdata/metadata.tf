resource "cloudflare_image" "%[1]s" {
  account_id = "%[2]s"
  id         = "%[1]s"
  url        = "%[3]s"
  creator    = "cftftest-creator"
  metadata = jsonencode({
    env    = "test"
    source = "terraform"
  })
}
