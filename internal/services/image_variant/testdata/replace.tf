resource "cloudflare_image_variant" "%[1]s" {
  account_id = "%[2]s"
  id         = "%[3]s"
  options = {
    fit      = "scale-down"
    width    = 100
    height   = 100
    metadata = "none"
  }
}
