package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareAddressMapSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description: "The account identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"description": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Description of the address map.",
		},
		"default_sni": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "If you have legacy TLS clients which do not send the TLS server name indicator, then you can specify one default SNI on the map.",
		},
		"enabled": {
			Type:        schema.TypeBool,
			Required:    true,
			Description: "Whether the Address Map is enabled or not.",
		},
		"can_delete": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "If set to false, then the Address Map cannot be deleted via API. This is true for Cloudflare-managed maps.",
		},
		"can_modify_ips": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "If set to false, then the IPs on the Address Map cannot be modified via the API. This is true for Cloudflare-managed maps.",
		},
		"ips": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "The set of IPs on the Address Map.",
			Elem: &schema.Resource{
				SchemaVersion: 1,
				Schema: map[string]*schema.Schema{
					"ip": {
						Type:         schema.TypeString,
						Required:     true,
						ValidateFunc: validation.IsIPAddress,
					},
				},
			},
		},
		"memberships": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "Zones and Accounts which will be assigned IPs on this Address Map.",
			Elem: &schema.Resource{
				SchemaVersion: 1,
				Schema: map[string]*schema.Schema{
					"identifier": {
						Type:     schema.TypeString,
						Required: true,
					},
					"kind": {
						Type:         schema.TypeString,
						Required:     true,
						ValidateFunc: validation.StringInSlice([]string{"account", "zone"}, false),
					},
					"can_delete": {
						Type:        schema.TypeBool,
						Computed:    true,
						Description: "Controls whether the membership can be deleted via the API or not.",
					},
				},
			},
		},
	}
}
