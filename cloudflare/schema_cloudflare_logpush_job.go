package cloudflare

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareLogpushJobSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Type:         schema.TypeString,
			Optional:     true,
			ExactlyOneOf: []string{"account_id", "zone_id"},
		},
		"zone_id": {
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
	}
}
