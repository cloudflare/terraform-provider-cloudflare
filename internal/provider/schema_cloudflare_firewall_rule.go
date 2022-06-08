package provider

import (
	"html"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareFirewallRuleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"zone_id": {
			Description: "The zone identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"filter_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"action": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice([]string{"block", "challenge", "allow", "js_challenge", "managed_challenge", "log", "bypass"}, false),
		},
		"priority": {
			Type:         schema.TypeInt,
			Optional:     true,
			ValidateFunc: validation.IntBetween(1, 2147483647),
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
		},
		"paused": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"products": {
			Type: schema.TypeSet,
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"zoneLockdown", "uaBlock", "bic", "hot", "securityLevel", "rateLimit", "waf"}, false),
			},
			Optional: true,
		},
	}
}
