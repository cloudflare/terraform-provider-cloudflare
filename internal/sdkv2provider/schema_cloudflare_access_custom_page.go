package sdkv2provider

import (
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareAccessCustomPageSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description:   consts.AccountIDSchemaDescription,
			Type:          schema.TypeString,
			Optional:      true,
			ForceNew:      true,
			ConflictsWith: []string{consts.ZoneIDSchemaKey},
		},
		consts.ZoneIDSchemaKey: {
			Description:   consts.ZoneIDSchemaDescription,
			Type:          schema.TypeString,
			Optional:      true,
			ForceNew:      true,
			ConflictsWith: []string{consts.AccountIDSchemaKey},
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Friendly name of the Access Custom Page configuration.",
		},
		"type": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice([]string{"identity_denied", "forbidden"}, false),
			Description:  fmt.Sprintf("Type of Access custom page to create. %s", renderAvailableDocumentationValuesStringSlice([]string{"identity_denied", "forbidden"})),
		},
		"custom_html": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Custom HTML to display on the custom page.",
		},
		"app_count": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Number of apps to display on the custom page.",
		},
	}
}
