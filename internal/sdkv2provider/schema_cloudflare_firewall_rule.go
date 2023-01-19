package sdkv2provider

import (
	"fmt"
	"html"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareFirewallRuleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.ZoneIDSchemaKey: {
			Description: "The zone identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"filter_id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The identifier of the Filter to use for determining if the Firewall Rule should be triggered.",
		},
		"action": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice([]string{"block", "challenge", "allow", "js_challenge", "managed_challenge", "log", "bypass"}, false),
			Description:  fmt.Sprintf("The action to apply to a matched request. %s", renderAvailableDocumentationValuesStringSlice([]string{"block", "challenge", "allow", "js_challenge", "managed_challenge", "log", "bypass"})),
		},
		"priority": {
			Type:         schema.TypeInt,
			Optional:     true,
			ValidateFunc: validation.IntBetween(1, 2147483647),
			Description:  "The priority of the rule to allow control of processing order. A lower number indicates high priority. If not provided, any rules with a priority will be sequenced before those without.",
		},
		"description": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringLenBetween(0, 500),
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				if html.UnescapeString(old) == html.UnescapeString(new) {
					return true
				}
				return false
			},
			Description: "A description of the rule to help identify it.",
		},
		"paused": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Whether this filter based firewall rule is currently paused.",
		},
		"products": {
			Type: schema.TypeSet,
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"zoneLockdown", "uaBlock", "bic", "hot", "securityLevel", "rateLimit", "waf"}, false),
			},
			Optional:    true,
			Description: fmt.Sprintf("List of products to bypass for a request when the bypass action is used. %s", renderAvailableDocumentationValuesStringSlice([]string{"zoneLockdown", "uaBlock", "bic", "hot", "securityLevel", "rateLimit", "waf"})),
		},
	}
}
