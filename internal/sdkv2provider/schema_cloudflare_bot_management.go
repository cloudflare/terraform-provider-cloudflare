package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareBotManagementSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.ZoneIDSchemaKey: {
			Description: consts.ZoneIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"enable_js": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Use lightweight, invisible JavaScript detections to improve Bot Management. [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).",
		},
		"fight_mode": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Whether to enable Bot Fight Mode.",
		},
		"sbfm_definitely_automated": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Super Bot Fight Mode (SBFM) action to take on definitely automated requests.",
		},
		"sbfm_likely_automated": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: " Super Bot Fight Mode (SBFM) action to take on likely automated requests.",
		},
		"sbfm_verified_bots": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Super Bot Fight Mode (SBFM) action to take on verified bots requests.",
		},
		"sbfm_static_resource_protection": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Super Bot Fight Mode (SBFM) to enable static resource protection. Enable if static resources on your application need bot protection. Note: Static resource protection can also result in legitimate traffic being blocked.",
		},
		"optimize_wordpress": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Whether to optimize Super Bot Fight Mode protections for Wordpress.",
		},
		"suppress_session_score": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Whether to disable tracking the highest bot score for a session in the Bot Management cookie.",
		},
		"auto_update_model": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Automatically update to the newest bot detection models created by Cloudflare as they are released. [Learn more.](https://developers.cloudflare.com/bots/reference/machine-learning-models#model-versions-and-release-notes)",
		},
		"using_latest_model": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "A read-only field that indicates whether the zone currently is running the latest ML model.",
		},
	}
}
