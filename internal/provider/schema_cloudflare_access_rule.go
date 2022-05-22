package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareAccessRuleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"zone_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
			Computed: true,
		},
		"mode": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice([]string{"block", "challenge", "whitelist", "js_challenge", "managed_challenge"}, false),
		},
		"notes": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"configuration": {
			Type:             schema.TypeList,
			MaxItems:         1,
			Required:         true,
			ForceNew:         true,
			DiffSuppressFunc: configurationDiffSuppress,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"target": {
						Type:         schema.TypeString,
						Required:     true,
						ForceNew:     true,
						ValidateFunc: validation.StringInSlice([]string{"ip", "ip6", "ip_range", "asn", "country"}, false),
					},
					"value": {
						Type:     schema.TypeString,
						Required: true,
						ForceNew: true,
					},
				},
			},
		},
	}
}
