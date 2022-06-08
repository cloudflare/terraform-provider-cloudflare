package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareWAFGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"group_id": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},

		"zone_id": {
			Description: "The zone identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},

		"package_id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},

		"mode": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "on",
			ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
		},
	}
}
