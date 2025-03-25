resource "cloudflare_zone_cache_variants" "%[2]s" {
	zone_id = "%[1]s"
	value = {
		avif = ["image/avif", "image/webp"]
		bmp = ["image/bmp", "image/webp"]
		gif = ["image/gif", "image/webp"]
		jpeg = ["image/jpeg", "image/webp"]
		jpg = ["image/jpg", "image/webp"]
		jp2 = ["image/jp2", "image/webp"]
		jpg2 = ["image/jpg2", "image/webp"]
		png = ["image/png"]
		tif = ["image/tif", "image/webp"]
		tiff = ["image/tiff", "image/webp"]
		webp = ["image/webp"]
	}
}
