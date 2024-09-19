package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareInfrastructureTargetSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description: consts.AccountIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
			Computed:    true,
		},
		"hostname": {
			Type:        schema.TypeString,
			Description: "A non-unique field that refers to a target.",
			Required:    true,
		},
		"ip": {
			Type:        schema.TypeList,
			Required:    true,
			Description: "The IPv4/IPv6 address that identifies where to reach a target.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"ipv4": {
						Type:        schema.TypeList,
						Optional:    true,
						Description: "The target's IPv4 address.",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"ip_addr": {
									Type:        schema.TypeString,
									Required:    true,
									Description: "The IP address of the target.",
								},
								"virtual_network_id": {
									Type:        schema.TypeString,
									Required:    true,
									Description: "The private virtual network identifier for the target.",
								},
							},
						},
					},
					"ipv6": {
						Type:        schema.TypeList,
						Optional:    true,
						Description: "The target's IPv6 address.",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"ip_addr": {
									Type:        schema.TypeString,
									Required:    true,
									Description: "The IP address of the target.",
								},
								"virtual_network_id": {
									Type:        schema.TypeString,
									Required:    true,
									Description: "The private virtual network identifier for the target.",
								},
							},
						},
					},
				},
			},
		},
		"created_at": {
			Type:     schema.TypeString,
			Optional: true,
			// Sets this value to read-only
			Computed:    true,
			Description: "The date and time at which the target was created.",
		},
		"modified_at": {
			Type:     schema.TypeString,
			Optional: true,
			// Sets this value to read-only
			Computed:    true,
			Description: "The date and time at which the target was modified.",
		},
	}
}
