package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareAccessBookmarkSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Description:   "The account identifier to target for the resource.",
			Type:          schema.TypeString,
			Optional:      true,
			Computed:      true,
			ConflictsWith: []string{"zone_id"},
		},
		"zone_id": {
			Description:   "The zone identifier to target for the resource.",
			Type:          schema.TypeString,
			Optional:      true,
			Computed:      true,
			ConflictsWith: []string{"account_id"},
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the bookmark application.",
		},
		"domain": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The domain of the bookmark application. Can include subdomains, paths, or both.",
		},
		"logo_url": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The image URL for the logo shown in the app launcher dashboard.",
		},
		"app_launcher_visible": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "Option to show/hide the bookmark in the app launcher.",
		},
	}
}
