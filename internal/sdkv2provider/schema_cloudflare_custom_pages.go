package sdkv2provider

import (
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareCustomPagesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.ZoneIDSchemaKey: {
			Description:   "The zone identifier to target for the resource.",
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{consts.AccountIDSchemaKey},
		},
		consts.AccountIDSchemaKey: {
			Description:   "The account identifier to target for the resource.",
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{consts.ZoneIDSchemaKey},
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
			Description: fmt.Sprintf("The type of custom page you wish to update. %s", renderAvailableDocumentationValuesStringSlice([]string{
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
			})),
		},
		"url": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "URL of where the custom page source is located.",
		},
		"state": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{"default", "customized"}, true),
			Description:  fmt.Sprintf("Managed state of the custom page. %s", renderAvailableDocumentationValuesStringSlice([]string{"default", "customized"})),
		},
	}
}
