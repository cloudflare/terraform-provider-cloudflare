package sdkv2provider

import (
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareTeamsListSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description: consts.AccountIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the teams list.",
		},
		"type": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice([]string{"IP", "SERIAL", "URL", "DOMAIN", "EMAIL"}, false),
			Description:  fmt.Sprintf("The teams list type. %s", renderAvailableDocumentationValuesStringSlice([]string{"IP", "SERIAL", "URL", "DOMAIN", "EMAIL"})),
		},
		"description": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The description of the teams list.",
		},
		"items": {
			Type:        schema.TypeSet,
			Optional:    true,
			MaxItems:    30000,
			Description: "The items of the teams list.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		// Adding items with description as optional separate field, so they
		// do not drown in between 1000s of string values at items attribute.
		// Use this field only if you have descriptions. The provider joins
		// items without description and this field together before processing.
		"items_with_description": {
			Type:        schema.TypeSet,
			Optional:    true,
			MaxItems:    30000,
			Description: "The items of the teams list that has explicit description.",
			ConfigMode:  schema.SchemaConfigModeAttr,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"value": {
						Type:     schema.TypeString,
						Required: true,
					},
					"description": {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
	}
}
