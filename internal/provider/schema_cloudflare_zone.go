package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareZoneSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"zone": {
			Type:             schema.TypeString,
			Required:         true,
			ForceNew:         true,
			DiffSuppressFunc: zoneDiffFunc,
		},
		"jump_start": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"paused": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"vanity_name_servers": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"plan": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			ValidateFunc: validation.StringInSlice([]string{
				planIDFree,
				planIDPro,
				planIDBusiness,
				planIDEnterprise,
				planIDPartnerFree,
				planIDPartnerPro,
				planIDPartnerBusiness,
				planIDPartnerEnterprise,
			}, false),
		},
		"meta": {
			Type:     schema.TypeMap,
			Computed: true,
			Elem: &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
		"status": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"type": {
			Type:         schema.TypeString,
			ValidateFunc: validation.StringInSlice([]string{"full", "partial"}, false),
			Default:      "full",
			Optional:     true,
		},
		"name_servers": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"verification_key": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}
