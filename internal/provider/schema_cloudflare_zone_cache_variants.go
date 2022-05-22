package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareZoneCacheVariantsExtensionSchema() *schema.Schema {
	return &schema.Schema{
		MinItems: 1,
		Optional: true,
		Type:     schema.TypeSet,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}
}

func resourceCloudflareZoneCacheVariantsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"zone_id": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"avif": resourceCloudflareZoneCacheVariantsExtensionSchema(),
		"bmp":  resourceCloudflareZoneCacheVariantsExtensionSchema(),
		"gif":  resourceCloudflareZoneCacheVariantsExtensionSchema(),
		"jpeg": resourceCloudflareZoneCacheVariantsExtensionSchema(),
		"jpg":  resourceCloudflareZoneCacheVariantsExtensionSchema(),
		"jpg2": resourceCloudflareZoneCacheVariantsExtensionSchema(),
		"jp2":  resourceCloudflareZoneCacheVariantsExtensionSchema(),
		"png":  resourceCloudflareZoneCacheVariantsExtensionSchema(),
		"tiff": resourceCloudflareZoneCacheVariantsExtensionSchema(),
		"tif":  resourceCloudflareZoneCacheVariantsExtensionSchema(),
		"webp": resourceCloudflareZoneCacheVariantsExtensionSchema(),
	}
}
