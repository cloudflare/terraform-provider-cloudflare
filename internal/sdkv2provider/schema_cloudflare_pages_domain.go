package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflarePagesDomainSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description: consts.AccountIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		// Domain name is the unique identifier for this resource, we use this in the API calls.
		// If this changes, `plan` will fail as it can't figure out the changes.
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
