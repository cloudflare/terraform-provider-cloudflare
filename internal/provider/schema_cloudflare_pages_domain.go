package provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func resourceCloudflarePagesDomainSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Description: "The account identifier to target for the resource",
			Type:        schema.TypeString,
			Required:    true,
		},
		"domain": {
			Description: "Name of the project",
			Type:        schema.TypeString,
			Required:    true,
		},
		"project_name": {
			Description: "Name of the project",
			Type:        schema.TypeString,
			Required:    true,
		},
		"id": {
			Description: "Id of the project",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"status": {
			Description: "Status of the custom domain",
			Type:        schema.TypeString,
			Computed:    true,
		},
	}
}
