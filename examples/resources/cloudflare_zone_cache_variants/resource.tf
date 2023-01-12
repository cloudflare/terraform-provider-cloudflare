resource "cloudflare_zone_cache_variants" "example" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  avif    = ["image/avif", "image/webp"]
  bmp     = ["image/bmp", "image/webp"]
  gif     = ["image/gif", "image/webp"]
  jpeg    = ["image/jpeg", "image/webp"]
  jpg     = ["image/jpg", "image/webp"]
  jpg2    = ["image/jpg2", "image/webp"]
  jp2     = ["image/jp2", "image/webp"]
  png     = ["image/png", "image/webp"]
  tiff    = ["image/tiff", "image/webp"]
  tif     = ["image/tif", "image/webp"]
  webp    = ["image/jpeg", "image/webp"]
}
