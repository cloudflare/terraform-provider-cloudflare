package cloudflare

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var supportedCacheVariantsExtensions = []string{
	"avif",
	"bmp",
	"gif",
	"jpeg",
	"jpg",
	"jpg2",
	"jp2",
	"png",
	"tiff",
	"tif",
	"webp",
}

func resourceCloudflareZoneCacheVariantsSchema() map[string]*schema.Schema {
	variantsSchema := map[string]*schema.Schema{
		"zone_id": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
	}

	for _, ext := range supportedCacheVariantsExtensions {
		variantsSchema[ext] = &schema.Schema{
			MinItems: 1,
			Optional: true,
			Type:     schema.TypeSet,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		}
	}

	return variantsSchema
}
