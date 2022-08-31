package provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func resourceCloudflareAPIShieldConfiguration() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"zone_id": {
			Description: "The zone identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"auth_id_characteristics": {
			Description: "Characteristics define properties across which auth-ids can be computed in a privacy-preserving manner.",
			Required:    true,
			Type:        schema.TypeList,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"type": {
						Description: "The type of characteristic. Currently accepts header or cookie.",
						Required:    true,
						Type:        schema.TypeString,
					},
					"name": {
						Description: "The name of the characteristic.",
						Required:    true,
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}
