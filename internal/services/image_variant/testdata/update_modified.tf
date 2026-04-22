resource "cloudflare_image_variant" "%[1]s" {
  account_id = "%[2]s"
  id         = "%[1]s"

  options = {
    fit      = "contain"
    height   = 1080
    width    = 1920
    metadata = "keep"
  }

  never_require_signed_urls = true
}
