package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareRegionalHostnameSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.ZoneIDSchemaKey: {
			Description: consts.ZoneIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
		},
		"hostname": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The hostname to regionalize.",
		},
		"region_key": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The region key. See [the full region list](https://developers.cloudflare.com/data-localization/regional-services/get-started/).",
		},
		"created_on": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The RFC3339 timestamp of when the hostname was created.",
		},
	}
}
