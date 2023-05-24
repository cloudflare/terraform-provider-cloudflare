package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareEmailRoutingSettingsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.ZoneIDSchemaKey: {
			Description: consts.ZoneIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"tag": {
			Description: "Email Routing settings identifier.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"name": {
			Description: "Domain of your zone.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"enabled": {
			Description: "State of the zone settings for Email Routing.",
			Type:        schema.TypeBool,
			Required:    true,
			ForceNew:    true,
		},
		"created": {
			Description: "The date and time the settings have been created.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"modified": {
			Description: "The date and time the settings have been modified.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"skip_wizard": {
			Description: "Flag to check if the user skipped the configuration wizard.",
			Type:        schema.TypeBool,
			Computed:    true,
			Optional:    true,
		},
		"status": {
			Description: "Show the state of your account, and the type or configuration error.",
			Type:        schema.TypeString,
			Computed:    true,
		},
	}
}
