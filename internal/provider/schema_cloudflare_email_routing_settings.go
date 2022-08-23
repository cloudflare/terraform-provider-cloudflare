package provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func resourceCloudflareEmailRoutingSettingsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"zone_id": {
			Description: "The zone identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
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
		},
	}
}
