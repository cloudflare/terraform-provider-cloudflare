package provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func resourceCloudflareAccessServiceTokenSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Description:   "The account identifier to target for the resource.",
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{"zone_id"},
		},
		"zone_id": {
			Description:   "The zone identifier to target for the resource.",
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{"account_id"},
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Friendly name of the token's intent.",
		},
		"client_id": {
			Type:        schema.TypeString,
			Computed:    true,
			ForceNew:    true,
			Description: "UUID client ID associated with the Service Token.",
		},
		"client_secret": {
			Type:        schema.TypeString,
			Computed:    true,
			Sensitive:   true,
			ForceNew:    true,
			Description: "A secret for interacting with Access protocols.",
		},
		"expires_at": {
			Type:        schema.TypeString,
			Computed:    true,
			ForceNew:    true,
			Description: "Date when the token expires",
		},
		"min_days_for_renewal": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     0,
			Description: "Regenerates the token if terraform is run within the specified amount of days before expiration",
		},
	}
}
