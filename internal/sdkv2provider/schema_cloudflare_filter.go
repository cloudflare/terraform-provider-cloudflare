package sdkv2provider

import (
	"html"
	"strings"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareFilterSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.ZoneIDSchemaKey: {
			Description: "The zone identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"paused": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Whether this filter is currently paused.",
		},
		"expression": {
			Type:     schema.TypeString,
			Required: true,
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				return strings.TrimSpace(new) == old
			},
			Description: "The filter expression to be used.",
		},
		"description": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringLenBetween(0, 500),
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				if html.UnescapeString(old) == html.UnescapeString(new) {
					return true
				}
				return false
			},
			Description: "A note that you can use to describe the purpose of the filter.",
		},
		"ref": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringLenBetween(0, 50),
			Description:  "Short reference tag to quickly select related rules.",
		},
	}
}
