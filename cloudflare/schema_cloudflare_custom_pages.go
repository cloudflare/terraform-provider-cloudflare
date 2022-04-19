package cloudflare

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareCustomPagesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"zone_id": {
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{"account_id"},
		},
		"account_id": {
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{"zone_id"},
		},
		"type": {
			Type:     schema.TypeString,
			Required: true,
			ValidateFunc: validation.StringInSlice([]string{
				"basic_challenge",
				"waf_challenge",
				"waf_block",
				"ratelimit_block",
				"country_challenge",
				"ip_block",
				"under_attack",
				"500_errors",
				"1000_errors",
				"always_online",
				"managed_challenge",
			}, true),
		},
		"url": {
			Type:     schema.TypeString,
			Required: true,
		},
		"state": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{"default", "customized"}, true),
		},
	}
}
