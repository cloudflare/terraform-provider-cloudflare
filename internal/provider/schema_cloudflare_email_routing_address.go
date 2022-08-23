package provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func resourceCloudflareEmailRoutingAddress() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Description: "The account identifier to target for the resource",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"tag": {
			Description: "Destination address identifier",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"email": {
			Description: "The contact email address of the user",
			Type:        schema.TypeString,
			Required:    true,
		},
		"verified": {
			Description: "The date and time the destination address has been verified. Null means not verified yet",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"created": {
			Description: "The date and time the destination address has been created.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"modified": {
			Description: "The date and time the destination address was last modified.\n",
			Type:        schema.TypeString,
			Computed:    true,
		},
	}
}
