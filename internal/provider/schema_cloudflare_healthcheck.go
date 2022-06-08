package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareHealthcheckSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"zone_id": {
			Description: "The zone identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"suspended": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"address": {
			Type:     schema.TypeString,
			Required: true,
		},
		"consecutive_fails": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
		},
		"consecutive_successes": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
		},
		"retries": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  2,
		},
		"timeout": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  5,
		},
		"interval": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  60,
		},
		"check_regions": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"WNAM", "ENAM", "WEU", "EEU", "NSAM", "SSAM", "OC", "ME", "NAF", "SAF", "IN", "SEAS", "NEAS", "ALL_REGIONS"}, false),
			},
		},
		"type": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice([]string{"TCP", "HTTP", "HTTPS"}, false),
		},
		"method": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validation.StringInSlice([]string{"connection_established", "GET", "HEAD"}, false),
		},
		"port": {
			Type:         schema.TypeInt,
			Optional:     true,
			Default:      80,
			ValidateFunc: validation.IntBetween(0, 65535),
		},
		"path": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "/",
		},
		"expected_codes": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"expected_body": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"follow_redirects": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"allow_insecure": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
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
		},
		"notification_suspended": {
			Type:       schema.TypeBool,
			Optional:   true,
			Default:    false,
			Deprecated: "Use `cloudflare_notification_policy` instead.",
		},
		"notification_email_addresses": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Deprecated: "Use `cloudflare_notification_policy` instead.",
		},
		"created_on": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"modified_on": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}
