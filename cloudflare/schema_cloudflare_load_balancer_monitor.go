package cloudflare

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareLoadBalancerMonitorSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"interval": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  60,
		},
		// interval has to be larger than (retries+1) * probe_timeout:

		"method": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},

		"port": {
			Type:         schema.TypeInt,
			Optional:     true,
			ValidateFunc: validation.IntBetween(0, 65535),
		},

		"retries": {
			Type:         schema.TypeInt,
			Optional:     true,
			Default:      2,
			ValidateFunc: validation.IntBetween(1, 5),
		},

		"timeout": {
			Type:         schema.TypeInt,
			Optional:     true,
			Default:      5,
			ValidateFunc: validation.IntBetween(1, 10),
		},

		"type": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "http",
			ValidateFunc: validation.StringInSlice([]string{"http", "https", "tcp", "udp_icmp", "icmp_ping", "smtp"}, false),
		},

		"created_on": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"modified_on": {
			Type:     schema.TypeString,
			Computed: true,
		},

		//
		// http/https only
		//
		"allow_insecure": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"expected_body": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"expected_codes": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"follow_redirects": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"header": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"header": {
						Type:     schema.TypeString,
						Required: true,
					},

					"values": {
						Type:     schema.TypeSet,
						Required: true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
				},
			},
			Set: HashByMapKey("header"),
		},

		"path": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},

		"probe_zone": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}
