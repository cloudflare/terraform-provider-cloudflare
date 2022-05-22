package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareBYOIPPrefixSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"prefix_id": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"description": {
			Type:     schema.TypeString,
			Computed: true,
			Optional: true,
		},
		"advertisement": {
			Type:         schema.TypeString,
			ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
			Computed:     true,
			Optional:     true,
		},
	}
}
