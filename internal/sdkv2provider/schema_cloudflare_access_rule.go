package sdkv2provider

import (
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareAccessRuleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description:  consts.AccountIDSchemaDescription,
			Type:         schema.TypeString,
			Optional:     true,
			ForceNew:     true,
			Computed:     true,
			ExactlyOneOf: []string{consts.AccountIDSchemaKey, consts.ZoneIDSchemaKey},
		},
		consts.ZoneIDSchemaKey: {
			Description:  consts.ZoneIDSchemaDescription,
			Type:         schema.TypeString,
			Optional:     true,
			ForceNew:     true,
			Computed:     true,
			ExactlyOneOf: []string{consts.AccountIDSchemaKey, consts.ZoneIDSchemaKey},
		},
		"mode": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice([]string{"block", "challenge", "whitelist", "js_challenge", "managed_challenge"}, false),
			Description:  fmt.Sprintf("The action to apply to a matched request. %s", renderAvailableDocumentationValuesStringSlice([]string{"block", "challenge", "whitelist", "js_challenge", "managed_challenge"})),
		},
		"notes": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "A personal note about the rule. Typically used as a reminder or explanation for the rule.",
		},
		"configuration": {
			Type:             schema.TypeList,
			MaxItems:         1,
			Required:         true,
			ForceNew:         true,
			DiffSuppressFunc: configurationDiffSuppress,
			Description:      "Rule configuration to apply to a matched request.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"target": {
						Type:         schema.TypeString,
						Required:     true,
						ForceNew:     true,
						ValidateFunc: validation.StringInSlice([]string{"ip", "ip6", "ip_range", "asn", "country"}, false),
						Description:  fmt.Sprintf("The request property to target. %s", renderAvailableDocumentationValuesStringSlice([]string{"ip", "ip6", "ip_range", "asn", "country"})),
					},
					"value": {
						Type:        schema.TypeString,
						Required:    true,
						ForceNew:    true,
						Description: "The value to target. Depends on target's type.",
					},
				},
			},
		},
	}
}
