// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bot_management

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func (r BotManagementResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"enable_js": schema.BoolAttribute{
				Description: "Use lightweight, invisible JavaScript detections to improve Bot Management. [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).",
				Optional:    true,
			},
			"fight_mode": schema.BoolAttribute{
				Description: "Whether to enable Bot Fight Mode.",
				Optional:    true,
			},
			"optimize_wordpress": schema.BoolAttribute{
				Description: "Whether to optimize Super Bot Fight Mode protections for Wordpress.",
				Optional:    true,
			},
			"sbfm_definitely_automated": schema.StringAttribute{
				Description: "Super Bot Fight Mode (SBFM) action to take on definitely automated requests.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("allow", "block", "managed_challenge"),
				},
			},
			"sbfm_static_resource_protection": schema.BoolAttribute{
				Description: "Super Bot Fight Mode (SBFM) to enable static resource protection.\nEnable if static resources on your application need bot protection.\nNote: Static resource protection can also result in legitimate traffic being blocked.\n",
				Optional:    true,
			},
			"sbfm_verified_bots": schema.StringAttribute{
				Description: "Super Bot Fight Mode (SBFM) action to take on verified bots requests.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("allow", "block"),
				},
			},
			"sbfm_likely_automated": schema.StringAttribute{
				Description: "Super Bot Fight Mode (SBFM) action to take on likely automated requests.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("allow", "block", "managed_challenge"),
				},
			},
			"auto_update_model": schema.BoolAttribute{
				Description: "Automatically update to the newest bot detection models created by Cloudflare as they are released. [Learn more.](https://developers.cloudflare.com/bots/reference/machine-learning-models#model-versions-and-release-notes)",
				Optional:    true,
			},
			"suppress_session_score": schema.BoolAttribute{
				Description: "Whether to disable tracking the highest bot score for a session in the Bot Management cookie.",
				Optional:    true,
			},
		},
	}
}
