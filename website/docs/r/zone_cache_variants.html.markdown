---
layout: "cloudflare"
page_title: "Cloudflare: cloudflare_zone_cache_variants"
sidebar_current: "docs-cloudflare-resource-zone-cache-variants"
description: |-
  Provides a resource which customizes Cloudflare zone cache variants setting.
---

# cloudflare_zone_cache_variants

Provides a resource which customizes Cloudflare zone cache variants.

## Example Usage

```hcl
resource "cloudflare_zone_cache_variants" "example" {
    zone_id = "7df50664b7f90274f4d77cdfee701380"

    avif = ["image/avif", "image/webp"]
    bmp = ["image/bmp", "image/webp"]
    gif = ["image/gif", "image/webp"]
    jpeg = ["image/jpeg", "image/webp"]
    jpg = ["image/jpg", "image/webp"]
    jpg2 = ["image/jpg2", "image/webp"]
    jp2 = ["image/jp2", "image/webp"]
    png = ["image/png", "image/webp"]
    tiff = ["image/tiff", "image/webp"]
    tif = ["image/tif", "image/webp"]
    webp = ["image/jpeg", "image/webp"]
}
```

## Argument Reference

The following arguments are supported:

* `zone_id` - (Required) The ID of the DNS zone in which to apply the cache variants setting
* `avif` - (Optional) List of strings with the MIME types of all the variants that should be served for avif
* `bmp` - (Optional) List of strings with the MIME types of all the variants that should be served for bmp
* `gif` - (Optional) List of strings with the MIME types of all the variants that should be served for gif
* `jpeg` - (Optional) List of strings with the MIME types of all the variants that should be served for jpeg
* `jpg` - (Optional) List of strings with the MIME types of all the variants that should be served for jpg
* `jpg2` - (Optional) List of strings with the MIME types of all the variants that should be served for jpg2
* `jp2` - (Optional) List of strings with the MIME types of all the variants that should be served for jp2
* `png` - (Optional) List of strings with the MIME types of all the variants that should be served for png
* `tiff` - (Optional) List of strings with the MIME types of all the variants that should be served for tiff
* `tif` - (Optional) List of strings with the MIME types of all the variants that should be served for tif
* `webp` - (Optional) List of strings with the MIME types of all the variants that should be served for webp

