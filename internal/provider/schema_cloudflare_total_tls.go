package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareTotalTLSSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"zone_id": {
			Description: "The zone identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"enabled": {
			Description: "Enable Total TLS for the zone.",
			Type:        schema.TypeBool,
			Required:    true,
		},
		"certificate_authority": {
			Description:  "The Certificate Authority that Total TLS certificates will be issued through.",
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{"google", "lets_encrypt"}, false),
		},
	}
}
