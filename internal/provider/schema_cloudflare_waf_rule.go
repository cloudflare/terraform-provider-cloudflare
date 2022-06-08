package provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func resourceCloudflareWAFRuleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"rule_id": {
			Type:     schema.TypeString,
			Required: true,
		},

		"zone_id": {
			Description: "The zone identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
		},

		"group_id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"package_id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},

		"mode": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}
