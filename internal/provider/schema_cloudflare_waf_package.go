package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareWAFPackageSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"package_id": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},

		"zone_id": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},

		"sensitivity": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "high",
			ValidateFunc: validation.StringInSlice([]string{"high", "medium", "low", "off"}, false),
		},

		"action_mode": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "challenge",
			ValidateFunc: validation.StringInSlice([]string{"simulate", "block", "challenge"}, false),
		},
	}
}
