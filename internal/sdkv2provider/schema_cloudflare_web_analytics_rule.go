package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareWebAnalyticsRuleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description: consts.AccountIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"ruleset_id": {
			Description: "The Web Analytics ruleset id.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
				diags := diag.Diagnostics{}
				if i.(string) == "" {
					diags = append(diags,
						diag.Diagnostic{
							Severity: diag.Error,
							Summary:  "The Web Analytics ruleset id is an empty string.",
							Detail:   "The provided Web Analytics ruleset id is empty. Verify the Web Analytics site has auto_install enabled.",
						})
				}
				return diags
			},
		},
		"host": {
			Description: "The host to apply the rule to.",
			Type:        schema.TypeString,
			Required:    true,
		},

		"paths": {
			Description: "A list of paths to apply the rule to.",
			Type:        schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Required: true,
		},

		"inclusive": {
			Description: "Whether the rule includes or excludes the matched traffic from being measured in Web Analytics.",
			Type:        schema.TypeBool,
			Required:    true,
		},

		"is_paused": {
			Description: "Whether the rule is paused or not.",
			Type:        schema.TypeBool,
			Required:    true,
		},
	}
}
