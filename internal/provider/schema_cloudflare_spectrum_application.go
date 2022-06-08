package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareSpectrumApplicationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"zone_id": {
			Description: "The zone identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
		},

		"protocol": {
			Type:     schema.TypeString,
			Required: true,
		},

		"traffic_type": {
			Type:     schema.TypeString,
			Default:  "direct",
			Optional: true,
			ValidateFunc: validation.StringInSlice([]string{
				"direct", "http", "https",
			}, false),
		},

		"dns": {
			Type:     schema.TypeList,
			Required: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"type": {
						Type:     schema.TypeString,
						Required: true,
					},
					"name": {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},

		"origin_direct": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},

		"origin_dns": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},

		"origin_port": {
			Type:          schema.TypeInt,
			Optional:      true,
			ConflictsWith: []string{"origin_port_range"},
			ValidateFunc:  validation.IntBetween(0, 65535),
		},

		"origin_port_range": {
			Type:          schema.TypeList,
			Optional:      true,
			MaxItems:      1,
			ConflictsWith: []string{"origin_port"},
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"start": {
						Type:         schema.TypeInt,
						Required:     true,
						ValidateFunc: validation.IntBetween(0, 65535),
					},
					"end": {
						Type:         schema.TypeInt,
						Required:     true,
						ValidateFunc: validation.IntBetween(0, 65535),
					},
				},
			},
		},

		"tls": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "off",
			ValidateFunc: validation.StringInSlice([]string{
				"off", "flexible", "full", "strict",
			}, false),
		},

		"ip_firewall": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},

		"proxy_protocol": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "off",
			ValidateFunc: validation.StringInSlice([]string{
				"off", "v1", "v2", "simple",
			}, false),
		},

		"edge_ips": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},

		"edge_ip_connectivity": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			ValidateFunc: validation.StringInSlice([]string{
				"all", "ipv4", "ipv6",
			}, false),
		},

		"argo_smart_routing": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
	}
}
