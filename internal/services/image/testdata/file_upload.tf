resource "cloudflare_image" "%[1]s" {
  account_id = "%[2]s"
  id         = "%[1]s"
  file       = "%[3]s"
  metadata = jsonencode({
    source = "terraform"
  })
}
