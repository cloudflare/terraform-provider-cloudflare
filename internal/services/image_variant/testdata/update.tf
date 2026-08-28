resource "cloudflare_image_variant" "%[1]s" {
  account_id = "%[2]s"
  id         = "%[1]s"
  options = {
    fit      = "cover"
    width    = 200
    height   = 200
    metadata = "copyright"
  }
  never_require_signed_urls = true
}
