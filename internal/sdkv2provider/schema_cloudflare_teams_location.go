package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var TeamsLocationNetworkSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"network": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "CIDR notation representation of the network IP.",
		},
	},
}

func resourceCloudflareTeamsLocationSchema() map[string]*schema.Schema {

	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description: consts.AccountIDSchemaDescription,
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
			ConfigMode:  schema.SchemaConfigModeAttr,
			Elem:        TeamsLocationNetworkSchema,
		},
		"client_default": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Indicator that this is the default location.",
		},
		"ecs_support": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Indicator that this location needs to resolve EDNS queries.",
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
			Description: "IPv4 to direct all IPv4 DNS queries to.",
		},
		"ipv4_destination_backup": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Backup IPv4 to direct all IPv4 DNS queries to.",
		},
		"dns_destination_ips_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "IPv4 binding assigned to this location",
		},
		"dns_destination_ipv6_block_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "IPv6 block binding assigned to this location",
		},
		"endpoints": {
			Type:        schema.TypeList,
			Description: "Endpoints assigned to this location",
			Optional:    true,
			MaxItems:    1,
			Elem:        TeamsLocationEndpointSchema,
		},
	}
}

var TeamsLocationEndpointSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"ipv4": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"enabled": {
						Type:     schema.TypeBool,
						Required: true,
					},
					"authentication_enabled": {
						Type:     schema.TypeBool,
						Computed: true,
					},
				},
			},
		},
		"ipv6": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"enabled": {
						Type:     schema.TypeBool,
						Required: true,
					},
					"authentication_enabled": {
						Type:     schema.TypeBool,
						Computed: true,
					},
					"networks": {
						Type:       schema.TypeList,
						ConfigMode: schema.SchemaConfigModeAttr,
						MinItems:   1,
						Optional:   true,
						Elem:       TeamsLocationNetworkSchema,
					},
				},
			},
		},
		"doh": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"require_token": {
						Type:     schema.TypeBool,
						Computed: true,
					},
					"enabled": {
						Type:     schema.TypeBool,
						Required: true,
					},
					"authentication_enabled": {
						Type:     schema.TypeBool,
						Computed: true,
					},
					"networks": {
						Type:       schema.TypeList,
						MinItems:   1,
						Optional:   true,
						Elem:       TeamsLocationNetworkSchema,
						ConfigMode: schema.SchemaConfigModeAttr,
					},
				},
			},
		},
		"dot": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"require_token": {
						Type:     schema.TypeBool,
						Computed: true,
					},
					"enabled": {
						Type:     schema.TypeBool,
						Required: true,
					},
					"authentication_enabled": {
						Type:     schema.TypeBool,
						Computed: true,
					},
					"networks": {
						Type:       schema.TypeList,
						Optional:   true,
						MinItems:   1,
						Elem:       TeamsLocationNetworkSchema,
						ConfigMode: schema.SchemaConfigModeAttr,
					},
				},
			},
		},
	},
}
