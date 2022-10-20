package provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func resourceCloudflarePagesDomainSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Description: "The account identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"domain": {
			Description: "Custom domain.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"project_name": {
			Description: "Name of the Pages Project.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"status": {
			Description: "Status of the custom domain.",
			Type:        schema.TypeString,
			Computed:    true,
		},
	}
}
