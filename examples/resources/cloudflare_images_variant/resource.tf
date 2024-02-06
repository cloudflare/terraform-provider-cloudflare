resource "cloudflare_images_variant" "example" {
  account_id                = "f037e56e89293a057740de681ac9abbe"
  variant_id                = "example"
  never_require_signed_urls = true
  options {
    fit      = "scale-down"
    metadata = "none"
    width    = 500
    height   = 500
  }
}
