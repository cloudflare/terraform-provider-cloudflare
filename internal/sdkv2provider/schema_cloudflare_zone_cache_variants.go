package sdkv2provider

import (
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareZoneCacheVariantsExtensionSchema(ext string) *schema.Schema {
	return &schema.Schema{
		MinItems: 1,
		Optional: true,
		Type:     schema.TypeSet,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Description: fmt.Sprintf("List of strings with the MIME types of all the variants that should be served for %s", ext),
	}
}

func resourceCloudflareZoneCacheVariantsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.ZoneIDSchemaKey: {
			Description: consts.ZoneIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"avif": resourceCloudflareZoneCacheVariantsExtensionSchema("avif"),
		"bmp":  resourceCloudflareZoneCacheVariantsExtensionSchema("bmp"),
		"gif":  resourceCloudflareZoneCacheVariantsExtensionSchema("gif"),
		"jpeg": resourceCloudflareZoneCacheVariantsExtensionSchema("jpeg"),
		"jpg":  resourceCloudflareZoneCacheVariantsExtensionSchema("jpg"),
		"jpg2": resourceCloudflareZoneCacheVariantsExtensionSchema("jpg2"),
		"jp2":  resourceCloudflareZoneCacheVariantsExtensionSchema("jp2"),
		"png":  resourceCloudflareZoneCacheVariantsExtensionSchema("png"),
		"tiff": resourceCloudflareZoneCacheVariantsExtensionSchema("tiff"),
		"tif":  resourceCloudflareZoneCacheVariantsExtensionSchema("tif"),
		"webp": resourceCloudflareZoneCacheVariantsExtensionSchema("webp"),
	}
}
