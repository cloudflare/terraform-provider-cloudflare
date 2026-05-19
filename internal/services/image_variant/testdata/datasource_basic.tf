resource "cloudflare_image_variant" "%[1]s" {
  account_id = "%[2]s"
  id         = "%[1]s"

  options = {
    fit      = "scale-down"
    height   = 480
    width    = 640
    metadata = "none"
  }
}

data "cloudflare_image_variant" "%[1]s" {
  account_id = "%[2]s"
  variant_id = cloudflare_image_variant.%[1]s.id
}
