package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareLogpushJobSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Description:  "The account identifier to target for the resource.",
			Type:         schema.TypeString,
			Optional:     true,
			ExactlyOneOf: []string{"account_id", "zone_id"},
		},
		"zone_id": {
			Description:  "The zone identifier to target for the resource.",
			Type:         schema.TypeString,
			Optional:     true,
			ExactlyOneOf: []string{"account_id", "zone_id"},
		},
		"enabled": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"dataset": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice([]string{"firewall_events", "http_requests", "spectrum_events", "nel_reports", "audit_logs", "gateway_dns", "gateway_http", "gateway_network", "dns_logs", "network_analytics_logs"}, false),
		},
		"logpull_options": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"destination_conf": {
			Type:     schema.TypeString,
			Required: true,
		},
		"ownership_challenge": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"filter": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"frequency": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "high",
			ValidateFunc: validation.StringInSlice([]string{"high", "low"}, false),
		},
	}
}
