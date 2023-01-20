package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareEmailRoutingAddressSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description: "The account identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"tag": {
			Description: "Destination address identifier.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"email": {
			Description: "The contact email address of the user.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"verified": {
			Description: "The date and time the destination address has been verified. Null means not verified yet.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"created": {
			Description: "The date and time the destination address has been created.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"modified": {
			Description: "The date and time the destination address was last modified.",
			Type:        schema.TypeString,
			Computed:    true,
		},
	}
}
