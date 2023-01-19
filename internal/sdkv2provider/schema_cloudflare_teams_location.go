package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareTeamsLocationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description: "The account identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the teams location.",
		},
		"networks": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "The networks CIDRs that comprise the location.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"id": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"network": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "CIDR notation representation of the network IP.",
					},
				},
			},
		},
		"client_default": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Indicator that this is the default location.",
		},
		"policy_ids": {
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Computed: true,
		},
		"ip": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Client IP address.",
		},
		"doh_subdomain": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The FQDN that DoH clients should be pointed at.",
		},
		"anonymized_logs_enabled": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Indicator that anonymized logs are enabled.",
		},
		"ipv4_destination": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "IP to direct all IPv4 DNS queries to.",
		},
	}
}
