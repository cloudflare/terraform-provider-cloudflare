package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareAccessBookmarkSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Type:          schema.TypeString,
			Optional:      true,
			Computed:      true,
			ConflictsWith: []string{"zone_id"},
		},
		"zone_id": {
			Type:          schema.TypeString,
			Optional:      true,
			Computed:      true,
			ConflictsWith: []string{"account_id"},
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"domain": {
			Type:     schema.TypeString,
			Required: true,
		},
		"logo_url": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"app_launcher_visible": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
	}
}
