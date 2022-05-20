package provider

import (
	"html"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareFilterSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"zone_id": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"paused": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"expression": {
			Type:     schema.TypeString,
			Required: true,
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				return strings.TrimSpace(new) == old
			},
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
		},
		"ref": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringLenBetween(0, 50),
		},
	}
}
