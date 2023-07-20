package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareManagedHeadersSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.ZoneIDSchemaKey: {
			Description: consts.ZoneIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
		},
		"managed_request_headers": {
			Description: "The list of managed request headers",
			Type:        schema.TypeSet,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"id": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "Unique headers rule identifier.",
					},
					"enabled": {
						Type:        schema.TypeBool,
						Required:    true,
						Description: "Whether the headers rule is active.",
					},
				},
			},
		},
		"managed_response_headers": {
			Description: "The list of managed response headers",
			Type:        schema.TypeSet,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"id": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "Unique headers rule identifier.",
					},
					"enabled": {
						Type:        schema.TypeBool,
						Required:    true,
						Description: "Whether the headers rule is active.",
					},
				},
			},
		},
	}
}
