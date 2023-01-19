package sdkv2provider

import (
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareBYOIPPrefixSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description: "The account identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"prefix_id": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "The assigned Bring-Your-Own-IP prefix ID.",
		},
		"description": {
			Type:        schema.TypeString,
			Computed:    true,
			Optional:    true,
			Description: "Description of the BYO IP prefix.",
		},
		"advertisement": {
			Type:         schema.TypeString,
			ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
			Computed:     true,
			Optional:     true,
			Description:  fmt.Sprintf("Whether or not the prefix shall be announced. A prefix can be activated or deactivated once every 15 minutes (attempting more regular updates will trigger rate limiting). %s", renderAvailableDocumentationValuesStringSlice([]string{"on", "off"})),
		},
	}
}
