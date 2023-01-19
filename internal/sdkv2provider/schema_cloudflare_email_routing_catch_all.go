package sdkv2provider

import (
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareEmailRoutingCatchAllSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.ZoneIDSchemaKey: {
			Description: "The zone identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"tag": {
			Description: "Routing rule identifier.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"name": {
			Description: "Routing rule name.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"enabled": {
			Description: "Routing rule status.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"matcher": {
			Description: "Matching patterns to forward to your actions.",
			Type:        schema.TypeSet,
			Required:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"type": {
						Description:  fmt.Sprintf("Type of matcher. %s", renderAvailableDocumentationValuesStringSlice([]string{"all"})),
						Type:         schema.TypeString,
						Required:     true,
						ValidateFunc: validation.StringInSlice([]string{"all"}, true),
					},
				},
			},
		},
		"action": {
			Description: "List actions patterns.",
			Type:        schema.TypeSet,
			Required:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"type": {
						Description:  fmt.Sprintf("Type of supported action. %s", renderAvailableDocumentationValuesStringSlice([]string{"drop", "forward", "worker"})),
						Type:         schema.TypeString,
						Required:     true,
						ValidateFunc: validation.StringInSlice([]string{"drop", "forward", "worker"}, true),
					},
					"value": {
						Description: "A list with items in the following form.",
						Type:        schema.TypeList,
						Required:    true,
						Elem: &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 90),
						},
					},
				},
			},
		},
	}
}
